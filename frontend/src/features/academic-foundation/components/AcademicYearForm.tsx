/**
 * Academic Year Form Component
 * Form for creating and updating academic years
 */

import { useState } from 'react';
import { Box, Button, TextField, Dialog, DialogTitle, DialogContent, DialogActions } from '@mui/material';
import { CreateAcademicYearRequest, UpdateAcademicYearRequest } from '@/api/academic-foundation';

interface AcademicYearFormProps {
  open: boolean;
  onClose: () => void;
  onSubmit: (data: CreateAcademicYearRequest | UpdateAcademicYearRequest) => void;
  editingData?: CreateAcademicYearRequest | UpdateAcademicYearRequest;
  title?: string;
}

export const AcademicYearForm = ({ open, onClose, onSubmit, editingData, title = 'Academic Year' }: AcademicYearFormProps) => {
  const [formData, setFormData] = useState<CreateAcademicYearRequest | UpdateAcademicYearRequest>(
    editingData || {
      school_id: '',
      name: '',
      start_date: '',
      end_date: '',
    }
  );

  const handleChange = (field: keyof (CreateAcademicYearRequest | UpdateAcademicYearRequest)) => (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setFormData({ ...formData, [field]: event.target.value });
  };

  const handleSubmit = () => {
    onSubmit(formData);
    onClose();
  };

  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <DialogTitle>{title}</DialogTitle>
      <DialogContent>
        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, pt: 2 }}>
          <TextField
            label="Nama Tahun Ajaran"
            fullWidth
            value={formData.name || ''}
            onChange={handleChange('name')}
            required
          />
          {/* Only show school_id when creating, not updating */}
          {!editingData && (
            <TextField
              label="ID Sekolah"
              fullWidth
              value={(formData as CreateAcademicYearRequest).school_id || ''}
              onChange={(e) => setFormData({ ...formData, school_id: e.target.value })}
              required
            />
          )}
          <TextField
            label="Tanggal Mulai"
            type="date"
            fullWidth
            InputLabelProps={{ shrink: true }}
            value={formData.start_date || ''}
            onChange={handleChange('start_date')}
            required
          />
          <TextField
            label="Tanggal Selesai"
            type="date"
            fullWidth
            InputLabelProps={{ shrink: true }}
            value={formData.end_date || ''}
            onChange={handleChange('end_date')}
            required
          />
        </Box>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Batal</Button>
        <Button onClick={handleSubmit} variant="contained">
          Simpan
        </Button>
      </DialogActions>
    </Dialog>
  );
};