import { request } from '@/utils/request';

/**
 * 登录
 * @returns 
 */
export const login = (params = {}) => {
  return request({
    url: '/login',
    method: 'post',
    data: {
      ...params
    }
  })
}