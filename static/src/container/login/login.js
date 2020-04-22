/**
 * Created by Max on 2019-03-25.
 */
import React from 'react'
import './login.css'
import PropTypes from 'prop-types'
import Modal from '../modal/modal'
import '../../compontent/reply'
import { getStore } from '../../App'
import request from '../../action/request'
import CloseImg from '../../img/closeImg'
import defaultIcon from '../../resource/png/default.png'
import { isPC } from '../../utils/utils'

class Login extends Modal {
  constructor(props) {
    super(props)
    const state = {
      name: '',
      email: '',
      blog: '',
      icon: '',
      file: null,
      dispatch: getStore().dispatch,
    }
    this.state = Object.assign({}, this.state, state)
  }

  onClickSubmit = () => {
    const {
      dispatch, name, email, blog, file,
    } = this.state
    const { onLoginSuccess, onLoginFailure } = this.props
    const form = new FormData()
    form.append('name', name)
    form.append('file', file)
    form.append('email', email)
    form.append('blog', blog)
    dispatch(request('register', form, 'post')).then((res) => {
      const { errormsg, errorcode, result } = res
      if (errorcode === 0) {
        const storage = window.localStorage
        storage.uuid = result.uuid
        if (onLoginSuccess) {
          onLoginSuccess()
        }
      } else if (onLoginFailure) {
        onLoginFailure(errormsg)
      }
    })
  };

  onGetImage = (e) => {
    const file = e.target.files[0]
    if (file === undefined) {
      return
    }
    const { type } = file
    if (/^image\/\S+$/.test(type)) {
      const src = URL.createObjectURL(file)
      this.setState({
        icon: src,
        file,
      })
    }
  };

  changeName = (e) => {
    this.setState({
      name: e.target.value,
    })
  };

  changeEmail = (e) => {
    this.setState({
      email: e.target.value,
    })
  };

  changeBlog = (e) => {
    this.setState({
      blog: e.target.value,
    })
  };

  subComponent = () => {
    const { onClickCancel } = this.props
    const { icon } = this.state
    const chooseIcon = icon.length ? icon : defaultIcon
    return (
      <div
        className={`login-container ${
          isPC() ? 'login-container-pc' : 'login-container-phone'
        }`}
      >
        <button className="login-close" type="button" onClick={onClickCancel}>
          <CloseImg className="login-close-img" />
        </button>
        <label htmlFor="upload-photo">
          <img src={chooseIcon} alt="" className="login-choose" />
        </label>
        <input
          id="upload-photo"
          type="file"
          // name="file"
          // id="input_file"
          accept="image/jpeg,image/jpg,image/png"
          onChange={this.onGetImage}
        />
        <label className="login-label" htmlFor="register">
          <input
            className="login-input"
            placeholder="昵称(必填...)"
            value={this.state.name}
            type="text"
            name="name"
            onChange={this.changeName}
          />
          <input
            className="login-input"
            placeholder="邮箱(不校验哦...)"
            value={this.state.email}
            type="text"
            name="name"
            onChange={this.changeEmail}
          />
          <input
            className="login-input"
            placeholder="博客（py必备...）"
            value={this.state.blog}
            type="text"
            name="name"
            onChange={this.changeBlog}
          />
        </label>
        <button
          className="reply-button"
          type="button"
          onClick={this.onClickSubmit}
        >
          提交
        </button>
      </div>
    )
  };
}

Login.PropTypes = {
  onLoginSuccess: PropTypes.func,
  onLoginFailure: PropTypes.func,
  onClickCancel: PropTypes.func,
}

Login.defaultProps = {
  onLoginSuccess: null,
  onLoginFailure: null,
  onClickCancel: null,
}

export default Login
