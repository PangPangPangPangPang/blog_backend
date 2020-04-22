/**
 * Created by Max on 2019-03-25.
 */
import React from 'react'
import './modal.css'
import ReactDom from 'react-dom'
import PropTypes from 'prop-types'

class Modal extends React.Component {
  constructor(props) {
    super(props)
    const rect = document.body.getBoundingClientRect()
    this.state = {
      height: document.body.clientHeight,
      width: document.body.clientWidth,
      top: -rect.top,
    }
  }

  componentDidMount() {
    const { show } = this.props
    if (show) {
      document.body.style.overflow = 'hidden'
    }
  }

  componentWillUnmount() {
    document.body.style.overflow = 'auto'
  }

  subComponent = () => null;

  onClickSubmit = () => {
    const { clickSubmit } = this.props
    clickSubmit()
  };

  onClickCancel = () => {
    const { clickCancel } = this.props
    if (clickCancel != null) {
      clickCancel()
    }
  };

  render() {
    const { show } = this.props
    const { width, height, top } = this.state
    if (!show) {
      return null
    }
    // 该方法可以指定父组件
    return ReactDom.createPortal(
      <div className="modal" style={{ height, width, top }}>
        <div>{this.subComponent()}</div>
      </div>,
      document.body,
    )
  }
}

Modal.propTypes = {
  show: PropTypes.bool,
  clickSubmit: PropTypes.func,
  clickCancel: PropTypes.func,
}

Modal.defaultProps = {
  show: false,
  clickSubmit: null,
  clickCancel: null,
}

export default Modal
