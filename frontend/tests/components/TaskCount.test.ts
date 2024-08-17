import { it, expect } from 'vitest'
import { renderSuspended } from '@nuxt/test-utils/runtime'
import { TaskCount } from '#components'
import { screen } from '@testing-library/vue'
import { registerEndpoint } from '@nuxt/test-utils/runtime'

it('can render TaskCount', async () => {
  registerEndpoint('/tasks/count', {
    method: 'GET',
    handler: () => ({ count: 3 })
  })

  const component = await renderSuspended(TaskCount)
  expect(screen.getByText('3')).toBeDefined() // Fails for now.
})
