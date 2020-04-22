/**
 * Created by Max on 2019-03-02 15:22
 */
import React from 'react'
import PropTypes from 'prop-types'
import './friendlink.css'

const FriendLink = (props) => {
  const { list } = props
  const ret = []
  for (let i = 0; i < list.length; i += 1) {
    const item = list[i]
    ret.push(
      <div key={i}>
        <a className="friend-link-name" href={item.link} target="_blank">
          {item.name}
        </a>
        {item.description}
      </div>,
    )
  }

  return (
    <div className="friend-link-detail">
      <div className="friend-link-header">友链出售中...</div>
      {ret}
    </div>
  )
}

FriendLink.propTypes = {
  list: PropTypes.array, // eslint-disable-line react/forbid-prop-types
}

FriendLink.defaultProps = {
  list: [],
}

export default FriendLink
