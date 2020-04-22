/**
 * Created by Max on 04/03/2017.
 */
import React from 'react'
import marked from 'marked'
import PropTypes from 'prop-types'
import hl from 'highlight.js'
import { ToastContainer, toast } from 'react-toastify'
import { connect } from 'react-redux'
import '../../../node_modules/highlight.js/styles/xcode.css'
import request from '../../action/request'
import { replyComment, clearReplyComment } from '../../action/reply'
import updateCurrentArticle from '../../action/user'
import './article.css'
import { getStore } from '../../App'
import Loading from '../../compontent/loading'
import Footer from '../footer/footer'
import DevImage from '../../resource/jpg/splatoon.png'
import Comment from '../../compontent/comment'
import Login from '../login/login'
import Reply from '../../compontent/reply'
import 'react-toastify/dist/ReactToastify.css'
import { isPC } from '../../utils/utils'

const renderer = new marked.Renderer()

marked.setOptions({
  renderer,
  gfm: true,
  tables: true,
  breaks: true,
  pedantic: true,
  sanitize: true,
  smartLists: true,
  smartypants: true,
})

renderer.code = (code, language) => {
  const validLang = !!(language && hl.getLanguage(language))
  const highlighted = validLang ? hl.highlight(language, code).value : code
  return `<pre><code class="article-code hljs ${language}">${highlighted}</code></pre>`
}

renderer.image = (href, title, text) => {
  let size
  if (title) {
    size = title.split('x')
  }
  if (title) {
    if (size[1]) {
      size = `width=${size[0]}height=${size[1]}`
    } else {
      size = `width=${size[0]}`
    }
  } else {
    size = '100%'
  }
  // Use placeholder image in the development environment.
  let h = href
  if (process.env.NODE_ENV === 'development') {
    h = DevImage
  }

  return `<img onclick="window.open(this.src)" class="article-img" src="${h}" width="100%" height="auto" alt=${text}/>`
}

let articleObject
renderer.heading = function heading(text, level) {
  const font = 30 - level * 3
  if (level === 1) {
    const timeStr = `<div style="background-color: #D0104C;padding-left: 7px; padding-right: 7px; padding-bottom:3px; padding-top: 3px; margin-right: 10px; border-radius: 2px; color: white; font-size: 13px">
                      ${articleObject.time}
                    </div>`
    return `<div>
              <div class="article-head" style="font-size: ${font}px;margin-bottom: 10px;margin-top: 30px;padding-bottom: 8px">
                <strong>${text}</strong>
              </div>
              <div style="margin-bottom:10px;font-weight: bold;display: flex;">
                ${timeStr}
              </div>
            </div>`
  }
  return `<div class="article-head" style="font-size: ${font}px;margin-bottom: 10px;margin-top: 30px;padding-bottom: 8px">
            <strong>${text}</strong>
          </div>`
}

renderer.strong = function strong(text) {
  return `<strong>${text}</strong>`
}

renderer.paragraph = function paragraph(text) {
  return `<div class="article-paragraph">${text}</div>`
}

renderer.list = function list(body, ordered) {
  const t = ordered ? 'decimal' : 'disc'
  return `<ul style="list-style-type: ${t}; margin-top: 0.6em; font-size: 15px; margin-bottom: 1.3em;">${body}</ul>`
}

renderer.listitem = function listitem(text) {
  return `<li style="margin-bottom: 0.3em;">${text}</li>`
}

renderer.blockquote = function em(text) {
  return `<blockquote class='article-blockquote'>${text}</blockquote>`
}

renderer.codespan = function code(text) {
  return `<code class="article-codespan">${text}</code>`
}

class Article extends React.Component {
  constructor() {
    super()
    this.state = {
      login: null,
      pageClass: `article-page ${
        isPC() ? 'article-page-pc' : 'article-page-phone'
      }`,
    }
  }

  componentDidMount() {
    window.scrollTo(0, 0)
    const { dispatch, params } = this.props
    const dic = { id: params.id }
    const store = getStore()
    if (store.getState().request[params.id]) {
      return
    }
    const commentArgs = {}
    const articleID = params.id
    commentArgs.article_id = articleID
    dispatch(updateCurrentArticle(articleID))
    dispatch(request('article', dic, 'get'))
    dispatch(request('fetchcomment', commentArgs, 'get'))
  }

  onClickReply = (args) => {
    const { dispatch, isLoggedin } = this.props
    if (!isLoggedin) {
      this.setState({
        login: (
          <Login
            onLoginSuccess={this.onLoginSuccess}
            onClickCancel={this.onLoginCancel}
            onLoginFailure={this.onLoginFailure}
            show
          />
        ),
      })
    } else {
      dispatch(replyComment(args.commentId))
    }
  };

  onClickConfirm = (args) => {
    const { dispatch, isLoggedin } = this.props
    if (!isLoggedin) {
      this.setState({
        login: (
          <Login
            onLoginSuccess={this.onLoginSuccess}
            onClickCancel={this.onLoginCancel}
            onLoginFailure={this.onLoginFailure}
            show
          />
        ),
      })
      return
    }
    const { params } = this.props
    const articleID = params.id
    const { uuid } = window.localStorage
    const { reply, commentId } = args
    const form = new FormData()
    form.append('article_id', articleID)
    form.append('uuid', uuid)
    form.append('content', reply)
    form.append('parent_id', commentId)
    dispatch(request('addcomment', form, 'post'))
      .then((res) => {
        const { errormsg, errorcode } = res
        if (errorcode === 0) {
          const commentArgs = {}
          commentArgs.article_id = params.id
          dispatch(request('fetchcomment', commentArgs, 'get'))
          dispatch(clearReplyComment())
        } else {
          toast(errormsg, {
            position: toast.POSITION.BOTTOM_CENTER,
          })
        }
      })
      .catch(() => {})
  };

  getFooter = () => {
    const { displayLoading } = this.props
    if (displayLoading === 0) {
      return <Footer />
    }
    return null
  };

  getComments = () => {
    const root = this.generateCommentNodeTree()
    if (root === null) {
      return null
    }
    const list = []
    for (let i = 0; i < root.children.length; i += 1) {
      const node = root.children[i]
      const comment = this.generateCommentComponent(node)
      list.push(comment)
    }
    const { pageClass } = this.state
    return <div className={pageClass}>{list}</div>
  };

  generateCommentComponent = (node) => {
    const { commentID } = this.props
    const { content, name, blog } = node
    const iconUrl = node.icon_url
    const createDate = node.create_date
    const currentCommentID = node.comment_id
    const component = (
      <Comment
        key={currentCommentID}
        commentId={currentCommentID}
        content={content}
        name={name}
        iconUrl={iconUrl}
        showReply={commentID === currentCommentID}
        blog={blog}
        createDate={createDate}
        clickReply={this.onClickReply}
        subcomponent={node.children}
        onClickConfirm={this.onClickConfirm}
      />
    )
    return component
  };

  onLoginSuccess = () => {
    toast('愉快的py吧！', {
      position: toast.POSITION.BOTTOM_CENTER,
    })
    this.setState({
      login: null,
    })
  };

  onLoginCancel = () => {
    this.setState({
      login: null,
    })
  };

  onLoginFailure = (errormsg) => {
    toast(errormsg, {
      position: toast.POSITION.BOTTOM_CENTER,
    })
  };

  generateCommentNodeTree = () => {
    const { comments } = this.props
    if (comments === null || comments.length === 0) {
      return null
    }

    const root = { parent_id: -1, children: [] }

    const map = {}
    map[-1] = root
    Object.keys(comments).forEach((index) => {
      const comment = comments[index]
      comment.children = []
      const commentId = comment.comment_id
      if (commentId !== undefined) {
        map[commentId] = comment
      }
    })
    Object.keys(comments).forEach((index) => {
      const comment = comments[index]
      const parentId = comment.parent_id
      const parent = map[parentId]
      parent.children.push(comment)
    })
    return root
  };

  render() {
    const { args, displayLoading, content } = this.props
    const { login, pageClass } = this.state
    articleObject = args
    return (
      <div>
        <div className={pageClass}>
          <Loading show={displayLoading} />
          <div dangerouslySetInnerHTML={{ __html: marked(content) }} />
        </div>
        <div className={pageClass}>
          <Reply clickConfirm={this.onClickConfirm} />
        </div>
        {this.getComments()}
        {this.getFooter()}
        {login}
        <ToastContainer
          autoClose={1000}
          closeButton={<div />}
          hideProgressBar
        />
      </div>
    )
  }
}

Article.propTypes = {
  dispatch: PropTypes.func,
  content: PropTypes.string,
  params: PropTypes.object, // eslint-disable-line react/forbid-prop-types
  args: PropTypes.object, // eslint-disable-line react/forbid-prop-types
  displayLoading: PropTypes.number,
  comments: PropTypes.array, // eslint-disable-line react/forbid-prop-types
  commentID: PropTypes.number,
  isLoggedin: PropTypes.bool,
}

Article.defaultProps = {
  dispatch: {},
  params: {},
  args: {},
  content: '',
  displayLoading: 1,
  comments: [],
  commentID: -1,
  isLoggedin: false,
}

function mapStateToProps(state, ownProps) {
  const getShow = () => {
    const articleId = ownProps.params.id
    if (state.request[articleId]) {
      return 0
    }
    return 1
  }
  const getContent = () => {
    const articleId = ownProps.params.id
    if (state.request[articleId]) {
      return state.request[articleId].content
    }
    return ''
  }

  const getArgs = () => {
    const articleId = ownProps.params.id
    if (state.request[articleId]) {
      return state.request[articleId]
    }
    return {}
  }
  const getComments = () => {
    const articleId = ownProps.params.id
    if (state.request[articleId]) {
      return state.request[articleId].comments
    }
    return []
  }
  const getCommentID = () => {
    const { commentID } = state.reply
    if (commentID === undefined) {
      return -1
    }
    return commentID
  }

  // 获取登录状态
  const getLoginStatus = () => {
    const storage = window.localStorage
    // 获取storage中缓存的uuid
    const { uuid } = storage
    if (uuid === undefined || uuid.length === 0) {
      if (state.request.register === undefined) {
        return false
      }
      // 获取从登录response中获得的uuid
      const suuid = state.request.register.uuid
      if (suuid === undefined || suuid.length === 0) {
        return false
      }
      storage.uuid = suuid
    }
    return true
  }
  return {
    content: getContent(),
    displayLoading: getShow(),
    args: getArgs(),
    comments: getComments(),
    commentID: getCommentID(),
    isLoggedin: getLoginStatus(),
  }
}

export default connect(mapStateToProps)(Article)
