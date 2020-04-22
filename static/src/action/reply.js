/**
 * Created by Max on 2019-03-27.
 */

import Types from './type'

function clearReplyComment() {
  return {
    type: Types.CLEAR_REPLY_COMMENT,
  }
}

function replyComment(text) {
  return {
    type: Types.REPLY_COMMENT,
    commentID: text,
  }
}

export { clearReplyComment, replyComment }
