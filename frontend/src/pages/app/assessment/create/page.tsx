/**
 * Assessment Create Page - MIGRATED TO TANSTACK QUERY
 * Create new Assessment
 */

import React, { useState } from 'react';
import {
  Box,
  Typography,
  Button,
  Alert,
} from '@mui/material';
import {
  ArrowBack,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { useCreateAssessment } from '@/services/commands/AssessmentCommandService';
import { useAuth } from '@/features/auth';
import AssessmentForm from '@/components/assessment/AssessmentForm';

const AssessmentCreatePage: React.FC = () => {
  const navigate = useNavigate();
  const { user } = useAuth();
  const [error, setError] = useState<string | null>(null);

  // ✅ Using TanStack Query mutation instead of manual API call
  const createMutation = useCreateAssessment({
    onSuccess: (newAssessment) => {
      navigate(`/assessment/${newAssessment.id}`);
    },
    onError: (err) => {
      setError(err.message || 'Gagal membuat asesmen');
    },
  });

  const handleSubmit = async (values: any) => {
    if (!user?.id) return;
    setError(null);
    createMutation.mutate({ data: values, userId: user.id });
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
        loading={createMutation.isPending}
      />
    </Box>
  );
};

export default AssessmentCreatePage;
