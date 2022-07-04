import { moduleMetadata, Story, Meta } from '@storybook/angular';
import { ProjectOverviewComponent } from './project-overview.component';

export default {
  title: 'ProjectOverviewComponent',
  component: ProjectOverviewComponent,
  decorators: [
    moduleMetadata({
      imports: [],
    })
  ],
} as Meta<ProjectOverviewComponent>;

const Template: Story<ProjectOverviewComponent> = (args: ProjectOverviewComponent) => ({
  props: args,
});


export const Primary = Template.bind({});
Primary.args = {
}