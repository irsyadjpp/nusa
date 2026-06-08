/**
 * Evidence Review Component
 * Review component for evidence
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Chip,
  Button,
  Grid,
  Avatar,
  IconButton,
} from '@mui/material';
import {
  Visibility,
  Edit,
  Delete,
  Assignment,
  Person,
  CalendarToday,
} from '@mui/icons-material';

interface Evidence {
  id: string;
  student_id: string;
  student_name: string;
  assessment_id: string;
  assessment_type: string;
  user_id: string;
  user_name: string;
  evidence_type: string;
  status: string;
  evidence_data: any;
  teacher_notes?: string;
  created_at: string;
  updated_at: string;
}

interface EvidenceReviewProps {
  evidence: Evidence;
  onView?: (id: string) => void;
  onEdit?: (id: string) => void;
  onDelete?: (id: string) => void;
  loading?: boolean;
}

const EvidenceReview: React.FC<EvidenceReviewProps> = ({
  evidence,
  onView,
  onEdit,
  onDelete,
  loading = false,
}) => {
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

  const getEvidenceTypeLabel = (type: string): string => {
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
    <Card>
      <CardContent>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
            <Avatar>{evidence.student_name.charAt(0)}</Avatar>
            <Box>
              <Typography variant="h6">{evidence.student_name}</Typography>
              <Typography variant="body2" color="text.secondary">
                {evidence.assessment_type}
              </Typography>
            </Box>
          </Box>
          <Chip
            label={evidence.status}
            color={getStatusColor(evidence.status)}
            size="small"
          />
        </Box>

        <Grid container spacing={2} sx={{ mb: 2 }}>
          <Grid item xs={6} sm={4}>
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              <Assignment fontSize="small" color="action" />
              <Box>
                <Typography variant="caption" color="text.secondary">
                  Tipe Bukti
                </Typography>
                <Typography variant="body2">
                  {getEvidenceTypeLabel(evidence.evidence_type)}
                </Typography>
              </Box>
            </Box>
          </Grid>
          <Grid item xs={6} sm={4}>
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              <Person fontSize="small" color="action" />
              <Box>
                <Typography variant="caption" color="text.secondary">
                  Guru
                </Typography>
                <Typography variant="body2">{evidence.user_name}</Typography>
              </Box>
            </Box>
          </Grid>
          <Grid item xs={6} sm={4}>
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              <CalendarToday fontSize="small" color="action" />
              <Box>
                <Typography variant="caption" color="text.secondary">
                  Tanggal
                </Typography>
                <Typography variant="body2">
                  {new Date(evidence.created_at).toLocaleDateString('id-ID')}
                </Typography>
              </Box>
            </Box>
          </Grid>
        </Grid>

        {evidence.teacher_notes && (
          <Box sx={{ mb: 2, p: 2, bgcolor: 'grey.50', borderRadius: 1 }}>
            <Typography variant="caption" color="text.secondary" gutterBottom>
              Catatan Guru:
            </Typography>
            <Typography variant="body2">{evidence.teacher_notes}</Typography>
          </Box>
        )}

        <Box sx={{ display: 'flex', gap: 1, justifyContent: 'flex-end' }}>
          {onView && (
            <Button
              size="small"
              variant="outlined"
              startIcon={<Visibility />}
              onClick={() => onView(evidence.id)}
              disabled={loading}
            >
              Lihat
            </Button>
          )}
          {onEdit && (
            <Button
              size="small"
              variant="outlined"
              startIcon={<Edit />}
              onClick={() => onEdit(evidence.id)}
              disabled={loading}
            >
              Edit
            </Button>
          )}
          {onDelete && (
            <Button
              size="small"
              variant="outlined"
              color="error"
              startIcon={<Delete />}
              onClick={() => onDelete(evidence.id)}
              disabled={loading}
            >
              Hapus
            </Button>
          )}
        </Box>
      </CardContent>
    </Card>
  );
};

export default EvidenceReview;
