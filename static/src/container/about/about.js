/**
 * Created by Max on 02/03/2017.
 */
import React from 'react'
import './about.css'
import Footer from '../footer/footer'
import img1 from '../../resource/jpg/about1.jpg'
import img2 from '../../resource/jpg/about2.jpg'
import img3 from '../../resource/jpg/about3.jpg'

export default class About extends React.Component {
  // constructor(props) {
  // super(props)
  // }
  render() {
    return (
      <div>
        <div className="about-card">
          <div className="about-item">
            <img src={img1} alt="" className="about-img about-img-left" />
            <div className="about-text">
              <p style={{ fontWeight: 'bold' }}>字节跳动</p>
              <p>目前在字节跳动打酱油,疯狂做业务，画界面</p>
            </div>
          </div>
          <div className="about-item">
            <div className="about-text">
              <p style={{ fontWeight: 'bold' }}>百度</p>
              <p>手机百度团队,负责RN部分业务</p>
            </div>
            <img src={img2} alt="" className="about-img about-img-right" />
          </div>
          <div className="about-item">
            <img src={img3} alt="" className="about-img about-img-left" />
            <div className="about-text">
              <p style={{ fontWeight: 'bold' }}>新浪</p>
              <p>视频模块，工程模块化，工程效率相关</p>
            </div>
          </div>
        </div>
        <Footer />
      </div>
    )
  }
}
