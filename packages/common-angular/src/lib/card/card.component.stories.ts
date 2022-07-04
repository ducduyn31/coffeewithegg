import { moduleMetadata, Story, Meta } from '@storybook/angular';
import { CardComponent } from './card.component';

export default {
  title: 'Card',
  component: CardComponent,
  decorators: [
    moduleMetadata({
      imports: [],
    })
  ],
} as Meta<CardComponent>;

const Template: Story<CardComponent> = (args: CardComponent) => ({
  props: args,
  template: `
      <div style='display: grid; grid-row-gap: 20px;'>
        <cwe-card>
          <div>Lorem ipsum dolor sit amet, consectetur adipisicing elit. A aliquid atque ea eos ex fugit magni modi neque odit officiis quam qui quisquam quod recusandae sed sint, sunt temporibus vero!</div>
        </cwe-card>

        <cwe-card style='width: 200px; height: 300px;'>
          <div>Lorem ipsum dolor sit amet, consectetur adipisicing elit. A aliquid atque ea eos ex fugit magni modi neque odit officiis quam qui quisquam quod recusandae sed sint, sunt temporibus vero!</div>
        </cwe-card>
      </div>
    `,
});


export const Primary = Template.bind({});
Primary.args = {
}
