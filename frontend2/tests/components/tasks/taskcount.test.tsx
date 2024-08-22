import { describe, expect, it, vi, afterEach } from 'vitest'
import * as queries from '../../../src/lib/queries'
import { signal } from '@preact/signals'
import { cleanup, render, screen } from '@testing-library/preact';
import TaskCount from '../../../src/components/tasks/taskcount'

describe('TaskCount', () => {
    afterEach(() => {
        // Cleanup the document tree to not assert other's test rendered components.
        cleanup()
    })

    it('Should display a loading message', async () => {
        const spy = vi.spyOn(queries, 'useTaskCount')
        spy.mockImplementation(() => ({
            data: signal(undefined),
            error: signal(null),
            isLoading: signal(true),
            refresh: () => {}
        }))

        render(<TaskCount />)

        const span = await screen.findByText('Loading...')
        expect(span).toBeDefined()
        const notPresent = screen.queryByText('Failed to load!')
        expect(notPresent).toBeNull()
    })

    it('Should display an error message', async () => {
        const spy = vi.spyOn(queries, 'useTaskCount')
        spy.mockImplementation(() => ({
            data: signal(undefined),
            error: signal('oops'),
            isLoading: signal(false),
            refresh: () => {}
        }))

        render(<TaskCount />)

        const span = await screen.findByText('Failed to load!')
        expect(span).toBeDefined()
        const notPresent2 = screen.queryByText('Loading...')
        expect(notPresent2).toBeNull()
    })

    
    it('Should display a task count', async () => {
        const spy = vi.spyOn(queries, 'useTaskCount')
        spy.mockImplementation(() => ({
            data: signal(3),
            error: signal(null),
            isLoading: signal(false),
            refresh: () => {}
        }))

        render(<TaskCount />)

        const span = await screen.findByText('3')
        expect(span).toBeDefined()
        const notPresent = screen.queryByText('Failed to load!')
        expect(notPresent).toBeNull()
        const notPresent2 = screen.queryByText('Loading...')
        expect(notPresent2).toBeNull()
    })
})
