/**
 * Created by Max on 03/03/2017.
 */
import React from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import './list.css'
import Summary from '../../compontent/summary'
import request from '../../action/request'
import { getStore } from '../../App'
import Loading from '../../compontent/loading'
import Footer from '../footer/footer'
import '../footer/footer.css'
import FriendLink from '../../compontent/friendlink'
import { isPC } from '../../utils/utils'

const FriendLinks = [
  {
    name: '鸳鸯',
    link: 'https://blog.y01.me',
    description: 'stk大佬',
  },
  {
    name: 'ギャラ',
    link: 'https://gyara.moe/',
    description: '菜鸡不多说',
  },
  {
    name: 'Xixing',
    link: 'https://xixing.dev/',
    description: 'P11',
  },
  {
    name: 'Dan',
    link: 'https://blog.igaryhe.io/',
    description: '@DGN',
  },
]

class List extends React.Component {
  static propTypes = {
    displayLoading: PropTypes.number,
    dispatch: PropTypes.func,
    list: PropTypes.array,
  };

  static defaultProps = {
    displayLoading: 1,
    dispatch: {},
    list: [],
  };

  constructor(props) {
    super(props)
    this.showFooter = false
    this.state = {
      content: getStore().getState().request.list || {},
    }
  }

  componentDidMount() {
    const { dispatch } = this.props
    const state = getStore().getState()
    if (!state.request.list) {
      dispatch(request('list'))
    }
  }

  getlist = () => {
    const ret = []
    const { list } = this.props
    const sourceList = list || []
    if (sourceList instanceof Array) {
      for (let i = 0; i < sourceList.length; i += 1) {
        ret.push(
          <Summary
            key={i}
            tags={sourceList[i].tags}
            name={sourceList[i].title}
            id={sourceList[i].id}
            time={sourceList[i].time}
          />,
        )
      }
    }
    if (ret.length > 0) {
      this.showFooter = true
    }
    return ret
  };

  getFooter = () => {
    if (this.showFooter) {
      return <Footer key={1000} />
    }
    return null
  };

  render() {
    const { displayLoading } = this.props
    return (
      <div>
        <div
          className={
            isPC() ? 'list-base list-base-pc' : 'list-base list-base-phone'
          }
        >
          <div
            className={
              isPC()
                ? 'list-template list-template-pc'
                : 'list-template list-template-phone'
            }
          >
            {this.getlist()}
            <Loading show={displayLoading} />
          </div>
          <div className="friend-link">
            <FriendLink list={FriendLinks} />
          </div>
        </div>
        {this.getFooter()}
      </div>
    )
  }
}

function mapStateToProps(state) {
  const getShow = () => {
    if (state.request.articles) {
      return 0
    }
    return 1
  }

  const getList = () => {
    if (state.request.articles) {
      return state.request.articles.list
    }
    return []
  }
  return {
    list: getList(),
    displayLoading: getShow(),
  }
}

export default connect(mapStateToProps)(List)
