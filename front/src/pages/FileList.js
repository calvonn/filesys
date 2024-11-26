// src/pages/FileList.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';

const FileList = () => {
    const [files, setFiles] = useState([]);

    useEffect(() => {
        const fetchFiles = async () => {
            try {
                const response = await axios.get('http://127.0.0.1:8080/files');
                setFiles(response.data.files);
            } catch (error) {
                console.error('Error fetching files:', error);
            }
        };
        fetchFiles();
    }, []);

    const handleDownload = (filename) => {
        const downloadUrl = `http://127.0.0.1:8080/download/${filename}`;
        window.location.href = downloadUrl;
    };

    return (
        <div>
            <h2>File List</h2>
            <ul>
                {files.map(file => (
                    <li key={file}>
                        <span>{file}</span>
                        <button onClick={() => handleDownload(file)}>Download</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default FileList;
