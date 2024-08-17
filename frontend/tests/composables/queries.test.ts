import { vi, it, expect, describe } from 'vitest'
import { registerEndpoint } from '@nuxt/test-utils/runtime'
import { useTaskCountQuery } from '../../composables/queries'
import { createError } from 'h3'

const taskCountEndpointMock = vi.fn()
registerEndpoint('/tasks/count', {
    method: 'GET',
    handler: taskCountEndpointMock
  })

describe('useTaskCountQuery', () => {
    it('Returns the count from the body', async () => {
        taskCountEndpointMock.mockImplementation(() => ({ count: 3 }))

        const { status, cnt } = await useTaskCountQuery()

        expect(status.value).toEqual("success")
        expect(cnt.value).toEqual(3)
    })

    it('Returns an error when the requests fails', async () => {
        taskCountEndpointMock.mockImplementation(() => { throw createError({message: "oops", statusCode: 400}) })

        const { status } = await useTaskCountQuery()

        expect(status.value).toEqual("error")
    })
})
