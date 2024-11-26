// src/pages/FileDownload.js
import React, { useState } from 'react';

const FileDownload = () => {
    const [fileId, setFileId] = useState('');
    const [downloadLink, setDownloadLink] = useState(null);

    const handleDownload = () => {
        // 模拟生成下载链接
        setDownloadLink(`http://127.0.0.1：3000/${fileId}`);
    };

    return (
        <div>
            <h2>File Download</h2>
            <input
                type="text"
                value={fileId}
                onChange={(e) => setFileId(e.target.value)}
                placeholder="Enter file ID"
            />
            <button onClick={handleDownload}>Download</button>
            {downloadLink && (
                <div>
                    <a href={downloadLink} download>Click here to download</a>
                </div>
            )}
        </div>
    );
};

export default FileDownload;
