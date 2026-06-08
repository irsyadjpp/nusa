/**
 * Evidence Upload Component
 * File upload component using react-dropzone
 */

import React, { useState, useCallback } from 'react';
import { useDropzone } from 'react-dropzone';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Button,
  LinearProgress,
  Chip,
  Alert,
  AlertTitle,
  IconButton,
  List,
  ListItem,
  ListItemText,
  ListItemIcon,
} from '@mui/material';
import {
  CloudUpload,
  Delete,
  CheckCircle,
  Error,
} from '@mui/icons-material';

interface EvidenceUploadProps {
  onUpload: (files: File[]) => void;
  acceptedFileTypes?: string[];
  maxFileSize?: number;
  maxFiles?: number;
  disabled?: boolean;
}

const EvidenceUpload: React.FC<EvidenceUploadProps> = ({
  onUpload,
  acceptedFileTypes = ['image/*', 'application/pdf', '.doc', '.docx'],
  maxFileSize = 10 * 1024 * 1024, // 10MB
  maxFiles = 5,
  disabled = false,
}) => {
  const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);
  const [uploadProgress, setUploadProgress] = useState<{ [key: string]: number }>({});
  const [error, setError] = useState<string | null>(null);

  const onDrop = useCallback(
    (acceptedFiles: File[]) => {
      setError(null);

      if (acceptedFiles.length > maxFiles) {
        setError(`Maksimal ${maxFiles} file yang dapat diupload`);
        return;
      }

      const validFiles = acceptedFiles.filter((file) => file.size <= maxFileSize);
      if (validFiles.length !== acceptedFiles.length) {
        setError('Beberapa file melebihi ukuran maksimal (10MB) dan tidak diupload');
      }

      setUploadedFiles((prev) => [...prev, ...validFiles]);

      // Simulate upload progress
      validFiles.forEach((file) => {
        let progress = 0;
        const interval = setInterval(() => {
          progress += 10;
          setUploadProgress((prev) => ({
            ...prev,
            [file.name]: progress,
          }));

          if (progress >= 100) {
            clearInterval(interval);
          }
        }, 100);
      });

      onUpload(validFiles);
    },
    [maxFiles, maxFileSize, onUpload]
  );

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    accept: acceptedFileTypes,
    maxFiles,
    disabled,
  });

  const handleRemoveFile = (fileToRemove: File) => {
    setUploadedFiles((prev) => prev.filter((file) => file !== fileToRemove));
    setUploadProgress((prev) => {
      const newProgress = { ...prev };
      delete newProgress[fileToRemove.name];
      return newProgress;
    });
  };

  const formatFileSize = (bytes: number): string => {
    if (bytes === 0) return '0 Bytes';
    const k = 1024;
    const sizes = ['Bytes', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return Math.round(bytes / Math.pow(k, i) * 100) / 100 + ' ' + sizes[i];
  };

  return (
    <Box>
      <Card>
        <CardContent>
          <Typography variant="h6" gutterBottom>
            Upload Bukti Pembelajaran
          </Typography>

          <Box
            {...getRootProps()}
            sx={{
              border: '2px dashed',
              borderColor: isDragActive ? 'primary.main' : 'grey.300',
              borderRadius: 2,
              p: 4,
              textAlign: 'center',
              cursor: disabled ? 'not-allowed' : 'pointer',
              bgcolor: isDragActive ? 'action.hover' : 'transparent',
              transition: 'all 0.3s',
              '&:hover': !disabled ? {
                borderColor: 'primary.main',
                bgcolor: 'action.hover',
              } : {},
            }}
          >
            <input {...getInputProps()} />
            <CloudUpload sx={{ fontSize: 48, color: 'text.secondary', mb: 2 }} />
            <Typography variant="body1" gutterBottom>
              {isDragActive
                ? 'Drop file di sini...'
                : 'Drag & drop file di sini, atau klik untuk memilih file'}
            </Typography>
            <Typography variant="caption" color="text.secondary">
              Maksimal {maxFiles} file, {formatFileSize(maxFileSize)} per file
            </Typography>
            <Box sx={{ mt: 2, display: 'flex', gap: 1, justifyContent: 'center', flexWrap: 'wrap' }}>
              {acceptedFileTypes.map((type, index) => (
                <Chip key={index} label={type} size="small" variant="outlined" />
              ))}
            </Box>
          </Box>

          {error && (
            <Alert severity="error" sx={{ mt: 2 }}>
              <AlertTitle>Error</AlertTitle>
              {error}
            </Alert>
          )}

          {uploadedFiles.length > 0 && (
            <Box sx={{ mt: 3 }}>
              <Typography variant="subtitle2" gutterBottom>
                File yang Diupload ({uploadedFiles.length}/{maxFiles})
              </Typography>
              <List>
                {uploadedFiles.map((file, index) => (
                  <ListItem
                    key={index}
                    secondaryAction={
                      <IconButton
                        edge="end"
                        onClick={() => handleRemoveFile(file)}
                        disabled={disabled}
                      >
                        <Delete />
                      </IconButton>
                    }
                  >
                    <ListItemIcon>
                      {uploadProgress[file.name] === 100 ? (
                        <CheckCircle color="success" />
                      ) : (
                        <Error color="warning" />
                      )}
                    </ListItemIcon>
                    <ListItemText
                      primary={file.name}
                      secondary={formatFileSize(file.size)}
                    />
                    <Box sx={{ width: '30%', ml: 2 }}>
                      <LinearProgress
                        variant="determinate"
                        value={uploadProgress[file.name] || 0}
                      />
                    </Box>
                  </ListItem>
                ))}
              </List>
            </Box>
          )}
        </CardContent>
      </Card>
    </Box>
  );
};

export default EvidenceUpload;
