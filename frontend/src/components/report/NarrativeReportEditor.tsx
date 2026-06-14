/**
 * Narrative Report Editor Component
 * Rich text editor for narrative
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Button,
} from '@mui/material';
import ReactQuill from 'react-quill-new';
import 'react-quill-new/dist/quill.snow.css';

interface NarrativeReportEditorProps {
  value: string;
  onChange: (value: string) => void;
  disabled?: boolean;
  onSave?: () => void;
  onCancel?: () => void;
  loading?: boolean;
}

const NarrativeReportEditor: React.FC<NarrativeReportEditorProps> = ({
  value,
  onChange,
  disabled = false,
  onSave,
  onCancel,
  loading = false,
}) => {
  const modules = {
    toolbar: [
      [{ header: [1, 2, 3, 4, 5, 6, false] }],
      [{ font: [] }],
      [{ size: [] }],
      ['bold', 'italic', 'underline', 'strike'],
      [{ color: [] }, { background: [] }],
      [{ list: 'ordered' }, { list: 'bullet' }],
      [{ align: [] }],
      ['link', 'image'],
      ['clean'],
    ],
  };

  const formats = [
    'header',
    'font',
    'size',
    'bold',
    'italic',
    'underline',
    'strike',
    'color',
    'background',
    'list',
    'bullet',
    'align',
    'link',
    'image',
  ];

  return (
    <Card>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Editor Rapor Naratif
        </Typography>

        <Box sx={{ mb: 2 }}>
          <ReactQuill
            theme="snow"
            value={value}
            onChange={onChange}
            modules={modules}
            formats={formats}
            placeholder="Tulis rapor naratif di sini..."
            readOnly={disabled}
            style={{ minHeight: '400px' }}
          />
        </Box>

        <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
          {onCancel && (
            <Button
              variant="outlined"
              onClick={onCancel}
              disabled={disabled || loading}
            >
              Batal
            </Button>
          )}
          {onSave && (
            <Button
              variant="contained"
              onClick={onSave}
              disabled={disabled || loading}
            >
              {loading ? 'Menyimpan...' : 'Simpan'}
            </Button>
          )}
        </Box>
      </CardContent>
    </Card>
  );
};

export default NarrativeReportEditor;
