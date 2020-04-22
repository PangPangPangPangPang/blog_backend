import React from 'react'
import './rss.css'
import rssImg from '../img/rssImg'
import { getbaseUrl } from '../action/request'

const Rss = () => {
  const Img = rssImg
  const onClickRss = () => {
    window.open(`${getbaseUrl()}feed`)
  }
  return (
    <div
      className="rss-container"
      onClick={onClickRss}
      role="button"
      onKeyPress={null}
      tabIndex="0"
    >
      <Img className="rss-image" />
    </div>
  )
}

export default Rss
