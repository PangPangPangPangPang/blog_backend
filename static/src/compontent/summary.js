/**
 * Created by Max on 03/03/2017.
 */
import React from 'react'
import { hashHistory } from 'react-router'
import './summary.css'
import PropTypes from 'prop-types'

const Summary = (props) => {
  const clickDetail = () => {
    hashHistory.push(`list/${props.id}`)
  }
  const renderTag = () => {
    const arr = []
    for (let i = 0; i < props.tags.length; i += 1) {
      arr.push(
        <div key={`${i}`} className="summary-tag">
          {props.tags[i]}
        </div>,
      )
    }
    return arr
  }
  const keyPress = () => {}
  const { name } = props
  return (
    <div
      className="summary-card"
      onClick={clickDetail}
      role="button"
      onKeyPress={keyPress}
      tabIndex={0}
    >
      {name}
      <hr className="summary-seperate" />
      <div className="summary-tags">{renderTag()}</div>
    </div>
  )
}
Summary.propTypes = {
  name: PropTypes.string,
  tags: PropTypes.array, // eslint-disable-line react/forbid-prop-types
  id: PropTypes.string,
}
Summary.defaultProps = {
  name: '',
  tags: [],
  id: '',
}

export default Summary
