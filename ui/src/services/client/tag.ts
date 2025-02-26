import useSWR from 'swr';

import request from '@answer/utils/request';
import { isLogin } from '@answer/utils';
import type * as Type from '@answer/common/interface';

export const deleteTag = (id) => {
  return request.delete('/answer/api/v1/tag', {
    tag_id: id,
  });
};
export const modifyTag = (params) => {
  return request.put('/answer/api/v1/tag', params);
};

export const useQuerySynonymsTags = (tagId) => {
  const apiUrl = tagId ? `/answer/api/v1/tag/synonyms?tag_id=${tagId}` : '';
  return useSWR<Type.SynonymsTag[]>(apiUrl, request.instance.get);
};

export const saveSynonymsTags = (params) => {
  return request.put('/answer/api/v1/tag/synonym', params);
};

export const useFollowingTags = () => {
  let apiUrl = '';
  if (isLogin()) {
    apiUrl = '/answer/api/v1/tags/following';
  }
  const { data, error, mutate } = useSWR<any[]>(apiUrl, request.instance.get);
  return {
    data,
    isLoading: !data && !error,
    error,
    mutate,
  };
};

export const useTagInfo = ({ id = '', name = '' }) => {
  let apiUrl;
  if (id) {
    apiUrl = `/answer/api/v1/tag/?id=${id}`;
  } else if (name) {
    apiUrl = `/answer/api/v1/tag/?name=${name}`;
  }
  const { data, error } = useSWR<Type.TagInfo>(apiUrl, request.instance.get);
  return {
    data,
    isLoading: !data && !error,
    error,
  };
};

export const followTags = (params) => {
  return request.put('/answer/api/v1/follow/tags', params);
};
