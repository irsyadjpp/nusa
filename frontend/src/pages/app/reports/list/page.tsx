/**
 * Report List Page
 * List view for Narrative Reports
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Grid,
  Card,
  CardContent,
  Chip,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  Add,
  Search,
  FilterList,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';

interface Report {
  id: string;
  student_id: string;
  student_name: string;
  period_start: string;
  period_end: string;
  status: string;
  created_at: string;
  published_at?: string;
}

const ReportListPage: React.FC = () => {
  const navigate = useNavigate();
  const [reports, setReports] = useState<Report[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedStudent, setSelectedStudent] = useState<string>('');
  const [selectedStatus, setSelectedStatus] = useState<string>('');

  useEffect(() => {
    loadReports();
  }, [selectedStudent, selectedStatus]);

  const loadReports = async () => {
    setLoading(true);
    setError(null);
    try {
      const params: any = {};
      if (selectedStudent) params.student_id = selectedStudent;
      if (selectedStatus) params.status = selectedStatus;

      const response = await fetch(`/api/reports?${new URLSearchParams(params)}`);
      const data = await response.json();
      setReports(data);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data Rapor');
    } finally {
      setLoading(false);
    }
  };

  const filteredReports = reports.filter((report) =>
    report.student_name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const getStatusColor = (status: string): 'success' | 'warning' | 'error' | 'info' | 'default' => {
    switch (status) {
      case 'PUBLISHED':
        return 'success';
      case 'DRAFT':
        return 'info';
      case 'PENDING':
        return 'warning';
      default:
        return 'default';
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4">Rapor Naratif</Typography>
        <Button
          variant="contained"
          startIcon={<Add />}
          onClick={() => navigate('/reports/generate')}
        >
          Buat Rapor Baru
        </Button>
      </Box>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Grid container spacing={2}>
            <Grid size={{ xs: 12, sm: 4 }}>
              <TextField
                fullWidth
                label="Cari Rapor"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                InputProps={{
                  startAdornment: <Search sx={{ mr: 1, color: 'text.secondary' }} />,
                }}
              />
            </Grid>
            <Grid size={{ xs: 12, sm: 4 }}>
              <FormControl fullWidth>
                <InputLabel>Status</InputLabel>
                <Select
                  value={selectedStatus}
                  label="Status"
                  onChange={(e) => setSelectedStatus(e.target.value)}
                >
                  <MenuItem value="">Semua</MenuItem>
                  <MenuItem value="DRAFT">Draft</MenuItem>
                  <MenuItem value="PENDING">Pending</MenuItem>
                  <MenuItem value="PUBLISHED">Dipublikasikan</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid size={{ xs: 12, sm: 4 }}>
              <Button
                fullWidth
                variant="outlined"
                startIcon={<FilterList />}
                onClick={() => {
                  setSearchTerm('');
                  setSelectedStudent('');
                  setSelectedStatus('');
                }}
              >
                Reset Filter
              </Button>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      {loading ? (
        <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
          <CircularProgress />
        </Box>
      ) : (
        <Grid container spacing={3}>
          {filteredReports.map((report) => (
            <Grid size={{ xs: 12, sm: 6, md: 4 }} key={report.id}>
              <Card
                sx={{ cursor: 'pointer', height: '100%' }}
                onClick={() => navigate(`/reports/${report.id}`)}
              >
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    {report.student_name}
                  </Typography>
                  <Typography variant="body2" color="text.secondary" gutterBottom>
                    {new Date(report.period_start).toLocaleDateString('id-ID')} -{' '}
                    {new Date(report.period_end).toLocaleDateString('id-ID')}
                  </Typography>
                  <Box sx={{ display: 'flex', gap: 1, mt: 2 }}>
                    <Chip
                      label={report.status}
                      color={getStatusColor(report.status)}
                      size="small"
                    />
                  </Box>
                  <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mt: 1 }}>
                    Dibuat: {new Date(report.created_at).toLocaleDateString('id-ID')}
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      )}

      {!loading && filteredReports.length === 0 && (
        <Alert severity="info">
          Tidak ada rapor yang ditemukan
        </Alert>
      )}
    </Box>
  );
};

export default ReportListPage;
