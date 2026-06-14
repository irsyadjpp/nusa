/**
 * Assessment List Page - MIGRATED TO TANSTACK QUERY
 * List view for Assessments
 */

import React, { useState } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  Chip,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Alert,
} from '@mui/material';
import {
  Add,
  Search,
  FilterList,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { useAssessments } from '@/services/queries/AssessmentQueryService';
import { AssessmentType, AssessmentStatus } from '@/shared/types/domain';

const AssessmentListPage: React.FC = () => {
  const navigate = useNavigate();
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedTP, setSelectedTP] = useState<string>('');
  const [selectedType, setSelectedType] = useState<AssessmentType | ''>('');
  const [selectedStatus, setSelectedStatus] = useState<AssessmentStatus | ''>('');

  // ✅ Using TanStack Query hook instead of manual state management
  const {
    data: assessments = [],
    error
  } = useAssessments({
    tp_id: selectedTP || undefined,
    assessment_type: selectedType || undefined,
    status: selectedStatus || undefined,
  });

  const filteredAssessments = assessments.filter((assessment) =>
    assessment.assessment_items?.questions?.some(q => 
      q.question_text.toLowerCase().includes(searchTerm.toLowerCase())
    ) || false
  );

  const getStatusColor = (status: AssessmentStatus): 'success' | 'warning' | 'error' | 'info' | 'default' => {
    switch (status) {
      case 'APPROVED':
        return 'success';
      case 'UNDER_REVIEW':
        return 'warning';
      case 'REJECTED':
        return 'error';
      case 'DRAFT':
        return 'info';
      case 'ARCHIVED':
        return 'default';
      default:
        return 'default';
    }
  };

  const getStatusLabel = (status: AssessmentStatus): string => {
    switch (status) {
      case 'APPROVED': return 'Disetujui';
      case 'UNDER_REVIEW': return 'Dalam Review';
      case 'REJECTED': return 'Ditolak';
      case 'DRAFT': return 'Draft';
      case 'ARCHIVED': return 'Diarsipkan';
      default: return status;
    }
  };

  const getTypeLabel = (type: AssessmentType): string => {
    switch (type) {
      case 'FORMATIVE': return 'Formatif';
      case 'SUMMATIVE': return 'Sumatif';
      default: return type;
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
          <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2 }}>
            <Box sx={{ width: { xs: '100%', sm: '33.33%' } }}>
              <TextField
                fullWidth
                label="Cari Asesmen"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                InputProps={{
                  startAdornment: <Search sx={{ mr: 1, color: 'text.secondary' }} />,
                }}
              />
            </Box>
            <Box sx={{ width: { xs: '100%', sm: '25%' } }}>
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
            </Box>
            <Box sx={{ width: { xs: '100%', sm: '16.67%' } }}>
              <FormControl fullWidth>
                <InputLabel>Tipe</InputLabel>
                <Select
                  value={selectedType}
                  label="Tipe"
                  onChange={(e) => setSelectedType(e.target.value as AssessmentType | '')}
                >
                  <MenuItem value="">Semua</MenuItem>
                  <MenuItem value="FORMATIVE">Formatif</MenuItem>
                  <MenuItem value="SUMMATIVE">Sumatif</MenuItem>
                </Select>
              </FormControl>
            </Box>
            <Box sx={{ width: { xs: '100%', sm: '16.67%' } }}>
              <FormControl fullWidth>
                <InputLabel>Status</InputLabel>
                <Select
                  value={selectedStatus}
                  label="Status"
                  onChange={(e) => setSelectedStatus(e.target.value as AssessmentStatus | '')}
                >
                  <MenuItem value="">Semua</MenuItem>
                  <MenuItem value="DRAFT">Draft</MenuItem>
                  <MenuItem value="UNDER_REVIEW">Dalam Review</MenuItem>
                  <MenuItem value="APPROVED">Disetujui</MenuItem>
                  <MenuItem value="REJECTED">Ditolak</MenuItem>
                  <MenuItem value="ARCHIVED">Diarsipkan</MenuItem>
                </Select>
              </FormControl>
            </Box>
            <Box sx={{ width: { xs: '100%', sm: '8.33%' } }}>
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
            </Box>
          </Box>
        </CardContent>
      </Card>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {typeof error === 'string' ? error : 'Error loading assessments'}
        </Alert>
      )}

      <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
          {filteredAssessments.map((assessment) => (
            <Box sx={{ width: { xs: '100%', sm: '50%', md: '33.33%' } }} key={assessment.id}>
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
                      label={getStatusLabel(assessment.status)}
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
            </Box>
          ))}
        </Box>

      {filteredAssessments.length === 0 && (
        <Alert severity="info">
          Tidak ada asesmen yang ditemukan
        </Alert>
      )}
    </Box>
  );
};

export default AssessmentListPage;
