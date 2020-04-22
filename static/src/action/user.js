/**
 * Created by Max on 03/03/2017.
 */

import Types from './type'

export default function updateCurrentArticle(text) {
  return {
    type: Types.UPDATE_CURRENT_ARTICLE,
    articleID: text,
  }
}
