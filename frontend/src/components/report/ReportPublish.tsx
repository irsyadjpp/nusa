/**
 * Report Publish Component
 * Publishing workflow for reports
 */

import React, { useState } from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Alert,
  AlertTitle,
  CircularProgress,
} from '@mui/material';
import {
  Publish,
  CheckCircle,
  Warning,
} from '@mui/icons-material';

interface ReportPublishProps {
  reportId: string;
  studentName: string;
  onPublish: (reportId: string, publishData: any) => void;
  loading?: boolean;
}

const ReportPublish: React.FC<ReportPublishProps> = ({
  reportId,
  studentName,
  onPublish,
  loading = false,
}) => {
  const [open, setOpen] = useState(false);
  const [publishTo, setPublishTo] = useState<'PARENTS' | 'STUDENT' | 'BOTH'>('PARENTS');
  const [message, setMessage] = useState('');
  const [confirmed, setConfirmed] = useState(false);

  const handlePublish = () => {
    onPublish(reportId, {
      publish_to: publishTo,
      message,
    });
    setOpen(false);
  };

  return (
    <>
      <Button
        variant="contained"
        startIcon={<Publish />}
        onClick={() => setOpen(true)}
        disabled={loading}
      >
        Publikasikan Rapor
      </Button>

      <Dialog open={open} onClose={() => setOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Publikasikan Rapor</DialogTitle>
        <DialogContent>
          <Alert severity="info" sx={{ mb: 2 }}>
            <AlertTitle>Informasi</AlertTitle>
            Anda akan mempublikasikan rapor untuk <strong>{studentName}</strong>.
            Tindakan ini akan mengirim notifikasi ke penerima yang dipilih.
          </Alert>

          <Box sx={{ mt: 2 }}>
            <FormControl fullWidth sx={{ mb: 2 }}>
              <InputLabel>Publikasikan ke</InputLabel>
              <Select
                value={publishTo}
                label="Publikasikan ke"
                onChange={(e) => setPublishTo(e.target.value as any)}
              >
                <MenuItem value="PARENTS">Orang Tua</MenuItem>
                <MenuItem value="STUDENT">Siswa</MenuItem>
                <MenuItem value="BOTH">Orang Tua dan Siswa</MenuItem>
              </Select>
            </FormControl>

            <TextField
              fullWidth
              multiline
              rows={4}
              label="Pesan Tambahan (Opsional)"
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              placeholder="Tambahkan pesan untuk penerima..."
            />

            <Box sx={{ mt: 2 }}>
              <Alert severity="warning">
                <AlertTitle>Konfirmasi</AlertTitle>
                Pastikan semua informasi dalam rapor sudah benar sebelum dipublikasikan.
              </Alert>
            </Box>

            <Box sx={{ mt: 2, display: 'flex', alignItems: 'center', gap: 1 }}>
              <input
                type="checkbox"
                checked={confirmed}
                onChange={(e) => setConfirmed(e.target.checked)}
              />
              <Typography variant="body2">
                Saya sudah memeriksa dan mengonfirmasi bahwa informasi dalam rapor sudah benar
              </Typography>
            </Box>
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpen(false)}>Batal</Button>
          <Button
            variant="contained"
            onClick={handlePublish}
            disabled={!confirmed || loading}
            startIcon={loading ? <CircularProgress size={20} /> : <Publish />}
          >
            {loading ? 'Memproses...' : 'Publikasikan'}
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default ReportPublish;
