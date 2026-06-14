/**
 * Report Publish Page
 * Separate page for publishing narrative reports
 */

import React, { useState } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  TextField,
  Alert,
  Container,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Publish as PublishIcon,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';

const ReportPublishPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const [publishTo, setPublishTo] = useState<'PARENTS' | 'STUDENT' | 'BOTH'>('PARENTS');
  const [message, setMessage] = useState('');
  const [confirmed, setConfirmed] = useState(false);
  const [submitting, setSubmitting] = useState(false);

  const handlePublish = () => {
    setSubmitting(true);
    // This would call the actual publish API
    console.log('Publishing report:', { reportId: id, publishTo, message });
    
    // Simulate API call
    setTimeout(() => {
      setSubmitting(false);
      navigate(`/dashboard/reports/${id}`);
    }, 1000);
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate(`/dashboard/reports/${id}`)}
          sx={{ mb: 2 }}
        >
          Kembali ke Detail Laporan
        </Button>
        <Typography variant="h4" component="h1">
          Publish Laporan Naratif
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
            <Alert severity="info" sx={{ mb: 2 }}>
              Publikasikan laporan naratif ini agar dapat dilihat oleh orang tua atau siswa.
            </Alert>
            
            <FormControl fullWidth>
              <InputLabel>Publikasikan ke</InputLabel>
              <Select
                value={publishTo}
                label="Publikasikan ke"
                onChange={(e) => setPublishTo(e.target.value as 'PARENTS' | 'STUDENT' | 'BOTH')}
              >
                <MenuItem value="PARENTS">Orang Tua</MenuItem>
                <MenuItem value="STUDENT">Siswa</MenuItem>
                <MenuItem value="BOTH">Orang Tua dan Siswa</MenuItem>
              </Select>
            </FormControl>
            
            <TextField
              label="Pesan (Opsional)"
              fullWidth
              multiline
              rows={3}
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              placeholder="Tambahkan pesan untuk penerima laporan"
            />

            <Box sx={{ display: 'flex', gap: 2, alignItems: 'center' }}>
              <input
                type="checkbox"
                id="confirmed"
                checked={confirmed}
                onChange={(e) => setConfirmed(e.target.checked)}
              />
              <label htmlFor="confirmed">
                Saya telah memeriksa isi laporan dan siap untuk mempublikasikannya
              </label>
            </Box>

            <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end', mt: 2 }}>
              <Button
                variant="outlined"
                onClick={() => navigate(`/dashboard/reports/${id}`)}
              >
                Batal
              </Button>
              <Button
                variant="contained"
                startIcon={<PublishIcon />}
                onClick={handlePublish}
                disabled={!confirmed || submitting}
              >
                {submitting ? 'Memproses...' : 'Publish'}
              </Button>
            </Box>
          </Box>
        </CardContent>
      </Card>
    </Container>
  );
};

export default ReportPublishPage;
