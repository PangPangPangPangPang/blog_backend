/**
 * Created by Max on 2017-03-31 16:57
 */

import React from 'react'
import './normal_button.css'
import PropTypes from 'prop-types'
import homeImg from '../img/homeImg'
import articleImg from '../img/articleImg'
import aboutImg from '../img/aboutImg'
import defaultImg from '../img/defaultImg'
import weiboImg from '../img/weiboImg'
import githubImg from '../img/githubImg'
import pushpinImg from '../img/pushpinImg'

class NormalButton extends React.Component {
  static propTypes = {
    // default button event handler
    handleClick: PropTypes.func,
    // // default button image
    img: PropTypes.string,
    // // defalut button content
    title: PropTypes.string,
  };

  static defaultProps = {
    handleClick: () => {},
    img: '',
    title: '',
  };

  constructor(props) {
    super(props)
    this.state = {
      backgroundColor: '',
      fontColor: '',
      imgColor: '',
    }
  }

  componentDidMount() {
    document.addEventListener('scroll', this.scrollEvent, false)
  }

  getImg = () => {
    let Img = null
    const { img } = this.props
    const { imgColor } = this.state
    switch (img) {
      case 'home': {
        Img = homeImg
        break
      }
      case 'tag': {
        Img = pushpinImg
        break
      }
      case 'article': {
        Img = articleImg
        break
      }
      case 'about': {
        Img = aboutImg
        break
      }
      case 'github': {
        Img = githubImg
        break
      }
      case 'weibo': {
        Img = weiboImg
        break
      }
      default: {
        Img = defaultImg
        break
      }
    }
    return <Img className={`normal-button-image ${imgColor}`} />
  };

  scrollEvent = () => {
    if (document.documentElement.scrollTop > 90) {
      this.setState({
        backgroundColor: 'normal-button-animation',
        fontColor: 'normal-button-title-animation',
        imgColor: 'normal-button-image-animation',
      })
    } else if (document.documentElement.scrollTop < 70) {
      this.setState({
        backgroundColor: 'normal-button-animation-back',
        fontColor: 'normal-button-title-animation-back',
        imgColor: 'normal-button-image-animation-back',
      })
    }
  };

  render() {
    const { handleClick, title } = this.props
    const { backgroundColor, fontColor } = this.state
    return (
      <button
        type="button"
        className={`normal-button ${backgroundColor}`}
        onClick={handleClick}
      >
        {this.getImg()}
        <div type="button" className={`normal-button-title ${fontColor}`}>
          {title}
        </div>
      </button>
    )
  }
}
export default NormalButton
