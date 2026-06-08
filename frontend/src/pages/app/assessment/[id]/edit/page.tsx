/**
 * Assessment Edit Page
 * Edit existing Assessment
 */

import React, { useState, useEffect } from 'react';
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
import { getAssessmentById, updateAssessment } from '@/api/assessment';
import { Assessment } from '@/api/assessment';
import AssessmentForm from '@/components/assessment/AssessmentForm';

const AssessmentEditPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [assessment, setAssessment] = useState<Assessment | null>(null);
  const [loading, setLoading] = useState(true);
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (id) {
      loadAssessment(id);
    }
  }, [id]);

  const loadAssessment = async (assessmentId: string) => {
    setLoading(true);
    setError(null);
    try {
      const data = await getAssessmentById(assessmentId);
      setAssessment(data);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data Asesmen');
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (values: any) => {
    if (!id) return;
    setSaving(true);
    setError(null);
    try {
      await updateAssessment(id, values);
      navigate(`/assessment/${id}`);
    } catch (err: any) {
      setError(err.message || 'Gagal mengupdate asesmen');
    } finally {
      setSaving(false);
    }
  };

  if (loading) {
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
        loading={saving}
      />
    </Box>
  );
};

export default AssessmentEditPage;
