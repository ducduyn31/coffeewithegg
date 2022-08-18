import React from 'react'
import { ComponentStory, ComponentMeta } from '@storybook/react'
import { icons } from '@coffeewithegg/base'
import Icon from './index'

export default {
  component: Icon,
  title: 'Icon',
} as ComponentMeta<typeof Icon>

const Template: ComponentStory<typeof Icon> = () => (
  <div
    style={{
      display: 'grid',
      gridTemplateColumns: 'repeat(6, minmax(0, 1fr))',
    }}
  >
    {icons.map((name) => (
      <div
        style={{
          display: 'flex',
          flexDirection: 'row',
          justifyContent: 'center',
          alignItems: 'center',
          columnGap: '20px',
        }}
      >
        <Icon name={name} />
        <p>{name}</p>
      </div>
    ))}
  </div>
)

export const Primary = Template.bind({})
Primary.args = {}
