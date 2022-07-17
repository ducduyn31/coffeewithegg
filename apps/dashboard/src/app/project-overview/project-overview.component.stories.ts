import { Meta, moduleMetadata, Story } from '@storybook/angular'
import { ProjectOverviewComponent } from './project-overview.component'
import { CommonAngularModule } from '@coffeewithegg/common-angular'

export default {
  title: 'ProjectOverviewComponent',
  component: ProjectOverviewComponent,
  decorators: [
    moduleMetadata({
      imports: [CommonAngularModule],
    }),
  ],
} as Meta<ProjectOverviewComponent>

const Template: Story<ProjectOverviewComponent> = (
  args: ProjectOverviewComponent,
) => ({
  props: args,
  template: `
      <coffeewithegg-project-overview [project]="
        {
          background: null,
          description: 'a website that is awesome',
          name: 'An Awesome Website',
          technologyUsed: ['angular', 'react']
        }
      "></coffeewithegg-project-overview>
    `,
})

export const Primary = Template.bind({})
Primary.args = {}
