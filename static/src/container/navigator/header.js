/**
 * Created by Max on 2017-03-31 11:26
 */
import React from 'react'
import { hashHistory } from 'react-router'
import './header.css'
import NormalButton from '../../compontent/normal_button'
import Logo from '../../img/logoImg'
import { isPC } from '../../utils/utils'

class Header extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      backgroundColor: '',
      iconColor: '',
    }
  }

  componentDidMount() {
    document.addEventListener('scroll', this.scrollEvent, false)
  }

  scrollEvent = () => {
    if (document.documentElement.scrollTop > 90) {
      this.setState({
        backgroundColor: 'header-animation',
        iconColor: 'header-default-icon-animation',
      })
    } else if (document.documentElement.scrollTop < 70) {
      this.setState({
        backgroundColor: 'header-animation-back',
        iconColor: 'header-default-icon-animation-back',
      })
    }
  };

  clickHome = () => {
    hashHistory.push('home')
  };

  clickArticle = () => {
    hashHistory.push('list')
  };

  clickTags = () => {
    hashHistory.push('tag')
  };

  clickAbout = () => {
    hashHistory.push('about')
  };

  render() {
    const { backgroundColor, iconColor } = this.state
    return (
      <div className="header-base">
        <div className={`header-default ${backgroundColor}`}>
          <NormalButton title="首页" img="home" handleClick={this.clickHome} />
          <NormalButton
            title="文章"
            img="article"
            handleClick={this.clickArticle}
          />
          <NormalButton title="Flag" img="tag" handleClick={this.clickTags} />
          <NormalButton
            title="关于"
            img="about"
            handleClick={this.clickAbout}
          />
          {isPC() ? (
            <Logo className={`header-default-icon ${iconColor}`} />
          ) : null}
        </div>
      </div>
    )
  }
}

export default Header
