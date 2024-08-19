import { vi, it, expect, describe } from 'vitest'
import { renderSuspended, mockNuxtImport } from '@nuxt/test-utils/runtime'
import { TaskCount } from '#components'

const { useTaskCountQueryMock } = vi.hoisted(() => ({ useTaskCountQueryMock: vi.fn() }))
mockNuxtImport('useTaskCountQuery', () => useTaskCountQueryMock)

describe("TaskCount", () => {
  it('can render TaskCount when query not completed yet', async () => {
    useTaskCountQueryMock.mockImplementation(() => ({ status: 'pending', cnt: -1 }))

    const component = await renderSuspended(TaskCount)
    expect(component.getByTestId("TaskCount").textContent).toEqual('Loading...')
  })

  it('can render TaskCount when query completed and is successful', async () => {
    useTaskCountQueryMock.mockImplementation(() => ({ status: 'success', cnt: 3 }))

    const component = await renderSuspended(TaskCount)
    expect(component.getByTestId("TaskCount").textContent).toEqual('3')
  })

  it('can render TaskCount when query completed and has failed', async () => {
    useTaskCountQueryMock.mockImplementation(() => ({ status: 'error', cnt: -1 }))

    const component = await renderSuspended(TaskCount)
    expect(component.getByTestId("TaskCount").textContent).toEqual('Error!')
  })
})
