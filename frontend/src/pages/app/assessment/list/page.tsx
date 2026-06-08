/**
 * Assessment List Page
 * List view for Assessments
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
import { getAssessments } from '@/api/assessment';
import { Assessment } from '@/api/assessment';

const AssessmentListPage: React.FC = () => {
  const navigate = useNavigate();
  const [assessments, setAssessments] = useState<Assessment[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedTP, setSelectedTP] = useState<string>('');
  const [selectedType, setSelectedType] = useState<string>('');
  const [selectedStatus, setSelectedStatus] = useState<string>('');

  useEffect(() => {
    loadAssessments();
  }, [selectedTP, selectedType, selectedStatus]);

  const loadAssessments = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await getAssessments({
        tp_id: selectedTP || undefined,
        assessment_type: selectedType || undefined,
        status: selectedStatus || undefined,
      });
      setAssessments(data);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data Asesmen');
    } finally {
      setLoading(false);
    }
  };

  const filteredAssessments = assessments.filter((assessment) =>
    assessment.assessment_items?.toString().toLowerCase().includes(searchTerm.toLowerCase())
  );

  const getStatusColor = (status: string): 'success' | 'warning' | 'error' | 'info' | 'default' => {
    switch (status) {
      case 'APPROVED':
        return 'success';
      case 'PENDING':
        return 'warning';
      case 'REJECTED':
        return 'error';
      case 'DRAFT':
        return 'info';
      default:
        return 'default';
    }
  };

  const getTypeLabel = (type: string): string => {
    switch (type) {
      case 'FORMATIVE':
        return 'Formatif';
      case 'SUMMATIVE':
        return 'Sumatif';
      default:
        return type;
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4">Asesmen</Typography>
        <Button
          variant="contained"
          startIcon={<Add />}
          onClick={() => navigate('/assessment/create')}
        >
          Buat Asesmen Baru
        </Button>
      </Box>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Grid container spacing={2}>
            <Grid item xs={12} sm={4}>
              <TextField
                fullWidth
                label="Cari Asesmen"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                InputProps={{
                  startAdornment: <Search sx={{ mr: 1, color: 'text.secondary' }} />,
                }}
              />
            </Grid>
            <Grid item xs={12} sm={3}>
              <FormControl fullWidth>
                <InputLabel>TP</InputLabel>
                <Select
                  value={selectedTP}
                  label="TP"
                  onChange={(e) => setSelectedTP(e.target.value)}
                >
                  <MenuItem value="">Semua</MenuItem>
                  {/* TP options would be loaded dynamically */}
                </Select>
              </FormControl>
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
                  <MenuItem value="FORMATIVE">Formatif</MenuItem>
                  <MenuItem value="SUMMATIVE">Sumatif</MenuItem>
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
                  <MenuItem value="DRAFT">Draft</MenuItem>
                  <MenuItem value="PENDING">Pending</MenuItem>
                  <MenuItem value="APPROVED">Disetujui</MenuItem>
                  <MenuItem value="REJECTED">Ditolak</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} sm={1}>
              <Button
                fullWidth
                variant="outlined"
                startIcon={<FilterList />}
                onClick={() => {
                  setSearchTerm('');
                  setSelectedTP('');
                  setSelectedType('');
                  setSelectedStatus('');
                }}
              >
                Reset
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
          {filteredAssessments.map((assessment) => (
            <Grid item xs={12} sm={6} md={4} key={assessment.id}>
              <Card
                sx={{ cursor: 'pointer', height: '100%' }}
                onClick={() => navigate(`/assessment/${assessment.id}`)}
              >
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    {getTypeLabel(assessment.assessment_type)}
                  </Typography>
                  <Typography variant="body2" color="text.secondary" gutterBottom>
                    TP ID: {assessment.tp_id}
                  </Typography>
                  <Box sx={{ display: 'flex', gap: 1, mb: 2 }}>
                    <Chip
                      label={assessment.status}
                      color={getStatusColor(assessment.status)}
                      size="small"
                    />
                    <Chip
                      label={`v${assessment.version_no}`}
                      size="small"
                      variant="outlined"
                    />
                  </Box>
                  {assessment.ai_confidence_score && (
                    <Typography variant="caption" color="text.secondary">
                      AI Confidence: {assessment.ai_confidence_score.toFixed(2)}
                    </Typography>
                  )}
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      )}

      {!loading && filteredAssessments.length === 0 && (
        <Alert severity="info">
          Tidak ada asesmen yang ditemukan
        </Alert>
      )}
    </Box>
  );
};

export default AssessmentListPage;
