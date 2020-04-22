/**
 * Created by Max on 2019-03-25.
 */
import React from 'react'
import { connect } from 'react-redux'
import './comment.css'
import PropTypes from 'prop-types'
import { getbaseUrl } from '../action/request'
import Reply from './reply'
import './reply.css'
import { getStore } from '../App'
import { replyComment } from '../action/reply'
import { isPC, getLocalTime } from '../utils/utils'
import defaultIcon from '../resource/png/default.png'

class Comment extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      dispatch: getStore().dispatch,
      marginLeft: isPC() ? '0px' : '0px',
    }
  }

  clickIconUrl = () => {
    const { blog } = this.props
    if (blog === null || blog.length === 0) {
      return
    }
    let des = blog
    if (!blog.startsWith('http')) {
      des = `http://${des}`
    }
    window.location.href = des
  };

  onClickReply = () => {
    const { dispatch } = this.state
    const { commentId } = this.props
    dispatch(replyComment(commentId))
  };

  onClickConfirm = (e) => {
    const { onClickConfirm, commentId } = this.props
    if (onClickConfirm) {
      const args = {
        reply: e.reply,
        commentId,
      }
      onClickConfirm(args)
    }
  };

  replyView = () => {
    const { showReply, commentId } = this.props
    if (!showReply) {
      return (
        <button
          className="reply-button"
          onClick={this.onClickReply}
          type="button"
        >
          回复
        </button>
      )
    }
    return <Reply clickConfirm={this.onClickConfirm} commentId={commentId} />
  };

  generateSub = () => {
    const { subcomponent } = this.props
    if (subcomponent.length <= 0) {
      return null
    }
    const list = []
    for (let i = 0; i < subcomponent.length; i += 1) {
      const node = subcomponent[i]
      const comment = this.generateCommentComponent(node)
      list.push(comment)
    }
    return <div className="comment-children">{list}</div>
  };

  generateCommentComponent = (node) => {
    const { clickReply, onClickConfirm } = this.props
    const { content, name, blog } = node
    const iconUrl = node.icon_url
    const createDate = node.create_date
    const commentID = node.comment_id
    const currentCommentID = getStore().getState().reply.commentID
    const { marginLeft } = this.state
    const component = (
      <Comment
        key={commentID}
        commentId={commentID}
        content={content}
        name={name}
        iconUrl={iconUrl}
        showReply={commentID === currentCommentID}
        blog={blog}
        createDate={createDate}
        clickReply={clickReply}
        subcomponent={node.children}
        marginLeft={marginLeft}
        onClickConfirm={onClickConfirm}
      />
    )
    return component
  };

  render() {
    const {
      iconUrl,
      name,
      createDate,
      content,
      marginLeft,
      showReply,
    } = this.props
    const formatDate = getLocalTime(createDate, 'yyyy-MM-dd hh:mm:ss')
    const icon = iconUrl.length
      ? `${getbaseUrl()}avatar/${iconUrl}`
      : defaultIcon
    return (
      <div
        className={`comment-container ${
          isPC() ? 'comment-container-pc' : 'comment-container-phone'
        }`}
        style={{ marginLeft }}
      >
        <div>
          <div className="comment-info">
            <div
              className="comment-icon"
              onClick={this.clickIconUrl}
              role="presentation"
            >
              <img className="comment-img" src={icon} alt="" />
            </div>
            <div className="comment-name">{name}</div>
            <div className="comment-date">{formatDate}</div>
          </div>
          <div
            className="comment-content"
            style={{ paddingBottom: showReply ? '10px' : '0px' }}
          >
            {content}
          </div>
          {this.replyView()}
          <div />
        </div>
        {this.generateSub()}
      </div>
    )
  }
}

Comment.propTypes = {
  iconUrl: PropTypes.string,
  name: PropTypes.string,
  blog: PropTypes.string,
  content: PropTypes.string,
  createDate: PropTypes.string,
  clickReply: PropTypes.func,
  showReply: PropTypes.bool,
  commentId: PropTypes.number,
  subcomponent: PropTypes.array, // eslint-disable-line react/forbid-prop-types

  marginLeft: PropTypes.string,
  onClickConfirm: PropTypes.func,
}

Comment.defaultProps = {
  iconUrl: 'default.png',
  name: '',
  blog: '',
  content: '',
  createDate: '',
  clickReply: null,
  showReply: false,
  commentId: -1,
  subcomponent: [],
  marginLeft: '0px',
  onClickConfirm: null,
}

function mapStateToProps(state, ownProps) {
  const getCommentID = () => {
    const { commentID } = state.reply
    if (!commentID) {
      return ownProps.currentCommentID
    }
    return commentID
  }
  return {
    currentCommentID: getCommentID(),
  }
}

export default connect(mapStateToProps)(Comment)
