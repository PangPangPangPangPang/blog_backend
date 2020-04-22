/**
 * Created by Max on 02/03/2017.
 */
// navigator height 48px

import React from 'react'
import { hashHistory } from 'react-router'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import Header from './navigator/header'
import './navigator/header.css'
import './navigator.css'
import Rss from '../compontent/rss'

class Navigator extends React.Component {
  static protoTypes = {
    children: PropTypes.Object,
  };

  static defaultProps = {
    children: {},
  };

  state = {
    current: 'mail',
  };

  handleClick = (e) => {
    switch (e.key) {
      case 'smile':
        hashHistory.push('about')
        break
      case 'home':
        hashHistory.push('home')
        break
      case 'article':
        hashHistory.push('list')
        break
      case 'tag':
        hashHistory.push('tag')
        break
      default:
        break
    }
    this.setState({
      current: e.key,
    })
  };

  render() {
    return (
      <div className="navigator">
        <Header />
        <div className="header-placeholder" />
        {this.props.children}
        <Rss />
      </div>
    )
  }
}

export default connect()(Navigator)
