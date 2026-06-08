/**
 * Teacher Feedback Component
 * Rich text feedback input using react-quill
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
} from '@mui/material';
import ReactQuill from 'react-quill-new';
import 'react-quill-new/dist/quill.snow.css';

interface TeacherFeedbackProps {
  value: string;
  onChange: (value: string) => void;
  disabled?: boolean;
  placeholder?: string;
}

const TeacherFeedback: React.FC<TeacherFeedbackProps> = ({
  value,
  onChange,
  disabled = false,
  placeholder = 'Masukkan feedback guru...',
}) => {
  const modules = {
    toolbar: [
      [{ header: [1, 2, 3, false] }],
      ['bold', 'italic', 'underline', 'strike'],
      [{ list: 'ordered' }, { list: 'bullet' }],
      ['clean'],
    ],
  };

  const formats = [
    'header',
    'bold',
    'italic',
    'underline',
    'strike',
    'list',
    'bullet',
  ];

  return (
    <Box>
      <Typography variant="subtitle2" gutterBottom>
        Feedback Guru
      </Typography>
      <Card>
        <CardContent sx={{ p: 0 }}>
          <ReactQuill
            theme="snow"
            value={value}
            onChange={onChange}
            modules={modules}
            formats={formats}
            placeholder={placeholder}
            readOnly={disabled}
            style={{ minHeight: '200px' }}
          />
        </CardContent>
      </Card>
    </Box>
  );
};

export default TeacherFeedback;
