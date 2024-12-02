// src/pages/FileUploadPage.js

import React, { useState } from 'react';
import { uploadFile } from '../services/api';

function FileUploadPage() {
    const [selectedFile, setSelectedFile] = useState(null);
    const [message, setMessage] = useState('');

    const handleFileChange = (event) => {
        setSelectedFile(event.target.files[0]);
    };

    const handleFileUpload = () => {
        if (!selectedFile) return;

        uploadFile(selectedFile).then(response => {
            setMessage(response.data.message);
            // 上传成功后，可以在这里更新文件列表或者做其他处理
        }).catch(error => {
            console.error("Error uploading file:", error);
            setMessage("文件上传失败");
        });
    };

    return (
        <div>
            <h1>文件上传</h1>
            <input type="file" onChange={handleFileChange} />
            <button onClick={handleFileUpload}>上传</button>
            {message && <p>{message}</p>}
        </div>
    );
}

export default FileUploadPage;
