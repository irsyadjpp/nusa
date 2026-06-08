/**
 * Assessment Create Page
 * Create new Assessment
 */

import React, { useState } from 'react';
import {
  Box,
  Typography,
  Button,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  ArrowBack,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { createAssessment } from '@/api/assessment';
import AssessmentForm from '@/components/assessment/AssessmentForm';

const AssessmentCreatePage: React.FC = () => {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (values: any) => {
    setLoading(true);
    setError(null);
    try {
      const newAssessment = await createAssessment(values);
      navigate(`/assessment/${newAssessment.id}`);
    } catch (err: any) {
      setError(err.message || 'Gagal membuat asesmen');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Button
            variant="outlined"
            startIcon={<ArrowBack />}
            onClick={() => navigate('/assessment')}
          >
            Kembali
          </Button>
          <Typography variant="h4">Buat Asesmen Baru</Typography>
        </Box>
      </Box>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      <AssessmentForm
        onSubmit={handleSubmit}
        onCancel={() => navigate('/assessment')}
        loading={loading}
      />
    </Box>
  );
};

export default AssessmentCreatePage;
