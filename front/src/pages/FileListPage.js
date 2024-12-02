import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { getFiles, downloadFile } from '../services/api';

function FileListPage() {
    const [files, setFiles] = useState([]);

    // 加载文件列表
    useEffect(() => {
        getFiles().then(response => {
            setFiles(response.data);
        }).catch(error => {
            console.error("Error fetching files:", error);
        });
    }, []);

    // 处理文件下载
    const handleDownload = (fileId, fileName) => {
        downloadFile(fileId).then(response => {
            const url = window.URL.createObjectURL(new Blob([response.data]));
            const link = document.createElement('a');
            link.href = url;
            link.setAttribute('download', fileName); // 设置下载文件名
            document.body.appendChild(link);
            link.click();
            link.remove();
        }).catch(error => {
            console.error("Error downloading file:", error);
        });
    };

    return (
        <div>
            <h1>文件列表</h1>
            <ul>
                {files.map(file => (
                    <li key={file.id}>
                        {file.name}
                        <button onClick={() => handleDownload(file.id, file.name)}>下载</button>
                    </li>
                ))}
            </ul>
            <p></p>
            <Link to="/upload">上传新文件</Link>
        </div>
    );
}

export default FileListPage;
