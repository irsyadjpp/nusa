/**
 * Assessment Edit Page - MIGRATED TO TANSTACK QUERY
 * Edit existing Assessment
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
import { useNavigate, useParams } from 'react-router-dom';
import { useAssessment } from '@/services/queries/AssessmentQueryService';
import { useUpdateAssessment } from '@/services/commands/AssessmentCommandService';
import AssessmentForm from '@/components/assessment/AssessmentForm';

const AssessmentEditPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [error, setError] = useState<string | null>(null);

  // ✅ Using TanStack Query hooks instead of manual state management
  const { 
    data: assessment, 
    isLoading 
  } = useAssessment(id!);

  const updateMutation = useUpdateAssessment({
    onSuccess: () => {
      navigate(`/assessment/${id}`);
    },
    onError: (err) => {
      setError(err.message || 'Gagal mengupdate asesmen');
    },
  });

  const handleSubmit = async (values: any) => {
    if (!id) return;
    setError(null);
    updateMutation.mutate({ id, data: values });
  };

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error || !assessment) {
    return (
      <Alert severity="error">
        {error || 'Asesmen tidak ditemukan'}
      </Alert>
    );
  }

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Button
            variant="outlined"
            startIcon={<ArrowBack />}
            onClick={() => navigate(`/assessment/${id}`)}
          >
            Kembali
          </Button>
          <Typography variant="h4">Edit Asesmen</Typography>
        </Box>
      </Box>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      <AssessmentForm
        initialValues={{
          tp_id: assessment.tp_id,
          tp_version_no: assessment.tp_version_no,
          assessment_type: assessment.assessment_type,
          assessment_items: assessment.assessment_items,
          answer_key: assessment.answer_key,
          scoring_guidelines: assessment.scoring_guidelines,
        }}
        onSubmit={handleSubmit}
        onCancel={() => navigate(`/assessment/${id}`)}
        loading={updateMutation.isPending}
      />
    </Box>
  );
};

export default AssessmentEditPage;
