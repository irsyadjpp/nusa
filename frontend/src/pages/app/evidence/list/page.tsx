/**
 * Evidence List Page
 * List view for Evidence
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Grid,
  Card,
  CardContent,
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
import { getEvidences } from '@/api/evidence';
import { Evidence } from '@/api/evidence';
import EvidenceReview from '@/components/evidence/EvidenceReview';

const EvidenceListPage: React.FC = () => {
  const navigate = useNavigate();
  const [evidences, setEvidences] = useState<Evidence[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedStudent, setSelectedStudent] = useState<string>('');
  const [selectedAssessment, setSelectedAssessment] = useState<string>('');
  const [selectedType, setSelectedType] = useState<string>('');
  const [selectedStatus, setSelectedStatus] = useState<string>('');

  useEffect(() => {
    loadEvidences();
  }, [selectedStudent, selectedAssessment, selectedType, selectedStatus]);

  const loadEvidences = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await getEvidences({
        student_id: selectedStudent || undefined,
        assessment_id: selectedAssessment || undefined,
        evidence_type: selectedType || undefined,
        status: selectedStatus || undefined,
      });
      setEvidences(data);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data Bukti');
    } finally {
      setLoading(false);
    }
  };

  const filteredEvidences = evidences.filter((evidence) =>
    evidence.student_name?.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const getStatusColor = (status: string): 'success' | 'warning' | 'error' | 'info' | 'default' => {
    switch (status) {
      case 'EVALUATED':
        return 'success';
      case 'PENDING':
        return 'warning';
      case 'REJECTED':
        return 'error';
      case 'SUBMITTED':
        return 'info';
      default:
        return 'default';
    }
  };

  const getTypeLabel = (type: string): string => {
    switch (type) {
      case 'STUDENT_WORK':
        return 'Karya Siswa';
      case 'ASSESSMENT_RESULT':
        return 'Hasil Asesmen';
      case 'OBSERVATION':
        return 'Observasi';
      default:
        return type;
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4">Bukti Pembelajaran</Typography>
        <Button
          variant="contained"
          startIcon={<Add />}
          onClick={() => navigate('/evidence/upload')}
        >
          Upload Bukti Baru
        </Button>
      </Box>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Grid container spacing={2}>
            <Grid item xs={12} sm={4}>
              <TextField
                fullWidth
                label="Cari Bukti"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                InputProps={{
                  startAdornment: <Search sx={{ mr: 1, color: 'text.secondary' }} />,
                }}
              />
            </Grid>
            <Grid item xs={12} sm={2}>
              <FormControl fullWidth>
                <InputLabel>Tipe</InputLabel>
                <Select
                  value={selectedType}
                  label="Tipe"
                  onChange={(e) => setSelectedType(e.target.value)}
                >
                  <MenuItem value="">Semua</MenuItem>
                  <MenuItem value="STUDENT_WORK">Karya Siswa</MenuItem>
                  <MenuItem value="ASSESSMENT_RESULT">Hasil Asesmen</MenuItem>
                  <MenuItem value="OBSERVATION">Observasi</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} sm={2}>
              <FormControl fullWidth>
                <InputLabel>Status</InputLabel>
                <Select
                  value={selectedStatus}
                  label="Status"
                  onChange={(e) => setSelectedStatus(e.target.value)}
                >
                  <MenuItem value="">Semua</MenuItem>
                  <MenuItem value="SUBMITTED">Dikirim</MenuItem>
                  <MenuItem value="PENDING">Pending</MenuItem>
                  <MenuItem value="EVALUATED">Dievaluasi</MenuItem>
                  <MenuItem value="REJECTED">Ditolak</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} sm={4}>
              <Button
                fullWidth
                variant="outlined"
                startIcon={<FilterList />}
                onClick={() => {
                  setSearchTerm('');
                  setSelectedStudent('');
                  setSelectedAssessment('');
                  setSelectedType('');
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
          {filteredEvidences.map((evidence) => (
            <Grid item xs={12} sm={6} md={4} key={evidence.id}>
              <EvidenceReview
                evidence={evidence}
                onView={(id) => navigate(`/evidence/${id}`)}
                onEdit={(id) => navigate(`/evidence/${id}/edit`)}
                onDelete={(id) => console.log('Delete evidence:', id)}
              />
            </Grid>
          ))}
        </Grid>
      )}

      {!loading && filteredEvidences.length === 0 && (
        <Alert severity="info">
          Tidak ada bukti yang ditemukan
        </Alert>
      )}
    </Box>
  );
};

export default EvidenceListPage;
