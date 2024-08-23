import { describe, expect, it, vi, afterEach, beforeAll, afterAll, beforeEach } from 'vitest'
import { renderHook, waitFor } from '@testing-library/preact';
import { clearQueryCache, useTaskCount } from '../../src/lib/queries';
import { http, HttpResponse } from 'msw';
import { setupServer } from 'msw/node';
import { API_BASE_URL } from '../../src/lib/config';


const server = setupServer()

beforeAll(() => server.listen({ onUnhandledRequest: 'error' }))
afterAll(() => server.close())
beforeEach(() => clearQueryCache())
afterEach(() => server.resetHandlers())

describe('useTaskCount', () => {
    it('Returns the task count', async () => {
        server.use(http.get(`${API_BASE_URL}/tasks/count`, () => HttpResponse.json({count: 3})))

        const query = renderHook(() => useTaskCount()).result.current

        await waitFor(() => {
            expect(query.isLoading.value).toEqual(false)
            expect(query.error.value).toBeNull()
            expect(query.data.value).toEqual(3)
        })
    })

    it('Handles errors', async () => {
        server.use(http.get(`${API_BASE_URL}/tasks/count`, () => HttpResponse.json({ error: 'oops' }, { status: 400 })))

        const query = renderHook(() => useTaskCount()).result.current

        await waitFor(() => {
            expect(query.isLoading.value).toEqual(false)
            expect(query.data.value).toBeUndefined()
            expect(query.error.value).toEqual({ error: 'oops' })
        })
    })
})
