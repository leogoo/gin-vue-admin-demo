import { request } from '@/utils/request';

/**
 * 判断是否已经初始化用户数据库
 * @returns 
 */
export const checkDB = () => {
  return request({
    url: '/init/checkdb',
    method: 'post'
  })
}