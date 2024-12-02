import axios from 'axios';

const API_BASE_URL = 'http://192.168.3.10:8222'; // 后端服务的基础URL

// 上传文件方法
export const uploadFile = (file) => {
    const formData = new FormData();
    formData.append('file', file);
    return axios.post(`${API_BASE_URL}/upload`, formData);
};

// 获取文件列表方法
export const getFiles = () => {
    return axios.get(`${API_BASE_URL}/list`);
};

// 下载文件方法
export const downloadFile = (fileId) => {
    return axios({
        url: `${API_BASE_URL}/download/${fileId}`,
        method: 'GET',
        responseType: 'blob', // 需要将响应类型设置为blob
    });
};
