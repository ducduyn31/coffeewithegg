import { moduleMetadata, Story, Meta } from '@storybook/angular';
import { IconComponent } from './icon.component';
import icons from './icon.list';

export default {
  title: 'Icons',
  component: IconComponent,
  decorators: [
    moduleMetadata({
      imports: [],
    })
  ],
} as Meta<IconComponent>;

const Template: Story<IconComponent> = (args: IconComponent) => ({
  props: args,
  template: `
      <div style='display: grid;grid-template-columns: repeat(6, minmax(0, 1fr))'>
        ${icons.map((name) => `
          <div style='display: flex; flex-direction: row; justify-content: center; align-items: center; column-gap: 20px;'>
            <cwe-icon name='${name}'></cwe-icon>
            <p>${name}</p>
          </div>
        `)}
      </div>
    `,
});


export const Primary = Template.bind({});
Primary.args = {
}
