import { describe, expect, it, vi, afterEach, beforeAll, afterAll } from 'vitest'
import { renderHook, waitFor } from '@testing-library/preact';
import { useTaskCount } from '../../src/lib/queries';
import { http, HttpResponse } from 'msw';
import { setupServer } from 'msw/node';
import { API_BASE_URL } from '../../src/lib/config';

const apiTaskCount = vi.fn()
const server = setupServer(
    http.get(`${API_BASE_URL}/tasks/count`, apiTaskCount)
)

beforeAll(() => server.listen({ onUnhandledRequest: 'error' }))
afterAll(() => server.close())
afterEach(() => server.resetHandlers())

describe('useTaskCount', () => {
    it('Returns the task count', async () => {
        apiTaskCount.mockImplementation(() => HttpResponse.json({count: 3}))

        const { result } = renderHook(() => useTaskCount())
        const resp = result.current

        await waitFor(() => {
            expect(resp.isLoading.value).toEqual(false)
            expect(resp.error.value).toBeNull()
            expect(resp.data.value).toEqual(3)
        })
    })

    it('Handles errors', async () => {
        apiTaskCount.mockImplementation(() => HttpResponse.json({ error: 'oops' }, { status: 400 }))

        const { result } = renderHook(() => useTaskCount())
        const resp = result.current

        await waitFor(() => {
            expect(resp.isLoading.value).toEqual(false)
            expect(resp.error.value).toEqual({ error: 'oops' })
            expect(resp.data.value).toBeUndefined()
        })
    })
})
