import { Project, Technology } from '@coffeewithegg/data-access'

export const mockProject = (mockInput?: {
  id?: string
  key?: string
  name?: string
  technologies?: Technology[]
  description?: string
}): Project => ({
  id: mockInput?.id || '1',
  key: mockInput?.key || 'key',
  name: mockInput?.name || 'name',
  description: mockInput?.description || 'description',
  technologies: mockInput?.technologies || [mockTechnology()],
  __typename: 'Project',
})

export const mockTechnology = (mockInput?: {
  id?: string
  key?: string
  name?: string
  description?: string
}): Technology => ({
  id: mockInput?.id || 'f79a4c54-8b76-4972-a030-1781a0689219',
  key: mockInput?.key || 'angular',
  name: mockInput?.name || 'Angular',
  description: mockInput?.description || '',
  __typename: 'Technology',
})
