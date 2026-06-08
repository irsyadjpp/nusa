/**
 * Revision History Component
 * Timeline of evaluation revisions
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Timeline,
  TimelineItem,
  TimelineSeparator,
  TimelineConnector,
  TimelineContent,
  TimelineDot,
  Chip,
  Avatar,
} from '@mui/material';
import {
  Edit,
  Person,
  AccessTime,
} from '@mui/icons-material';

interface Evaluation {
  id: string;
  revision_no: number;
  student_id: string;
  student_name: string;
  user_id: string;
  user_name: string;
  total_score: number;
  max_score: number;
  performance_level: string;
  teacher_feedback?: string;
  evaluated_at: string;
  created_at: string;
  updated_at: string;
}

interface RevisionHistoryProps {
  evaluations: Evaluation[];
}

const RevisionHistory: React.FC<RevisionHistoryProps> = ({ evaluations }) => {
  const getPerformanceLevelColor = (level: string): 'success' | 'info' | 'warning' | 'error' => {
    switch (level) {
      case 'EXCELLENT':
        return 'success';
      case 'PROFICIENT':
        return 'info';
      case 'DEVELOPING':
        return 'warning';
      case 'BEGINNING':
        return 'error';
      default:
        return 'info';
    }
  };

  const getPerformanceLevelLabel = (level: string): string => {
    switch (level) {
      case 'EXCELLENT':
        return 'Sangat Baik';
      case 'PROFICIENT':
        return 'Baik';
      case 'DEVELOPING':
        return 'Sedang Berkembang';
      case 'BEGINNING':
        return 'Perlu Bimbingan';
      default:
        return level;
    }
  };

  const sortedEvaluations = [...evaluations].sort(
    (a, b) => b.revision_no - a.revision_no
  );

  return (
    <Card>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Riwayat Revisi Evaluasi
        </Typography>

        {sortedEvaluations.length === 0 ? (
          <Typography variant="body2" color="text.secondary">
            Belum ada riwayat revisi
          </Typography>
        ) : (
          <Timeline>
            {sortedEvaluations.map((evaluation, index) => (
              <TimelineItem key={evaluation.id}>
                <TimelineSeparator>
                  <TimelineDot color={index === 0 ? 'primary' : 'grey'}>
                    <Edit fontSize="small" />
                  </TimelineDot>
                  {index < sortedEvaluations.length - 1 && <TimelineConnector />}
                </TimelineSeparator>
                <TimelineContent>
                  <Box
                    sx={{
                      p: 2,
                      border: 1,
                      borderColor: 'divider',
                      borderRadius: 1,
                      bgcolor: index === 0 ? 'primary.50' : 'background.paper',
                    }}
                  >
                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 2, mb: 1 }}>
                      <Avatar sx={{ width: 32, height: 32 }}>
                        {evaluation.user_name.charAt(0)}
                      </Avatar>
                      <Box sx={{ flexGrow: 1 }}>
                        <Typography variant="subtitle2" fontWeight="bold">
                          Revisi #{evaluation.revision_no}
                        </Typography>
                        <Typography variant="caption" color="text.secondary">
                          oleh {evaluation.user_name}
                        </Typography>
                      </Box>
                      <Chip
                        label={getPerformanceLevelLabel(evaluation.performance_level)}
                        color={getPerformanceLevelColor(evaluation.performance_level)}
                        size="small"
                      />
                    </Box>

                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 1, mb: 1 }}>
                      <AccessTime fontSize="small" color="action" />
                      <Typography variant="caption" color="text.secondary">
                        {new Date(evaluation.evaluated_at).toLocaleString('id-ID')}
                      </Typography>
                    </Box>

                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 2, mb: 1 }}>
                      <Typography variant="body2">
                        Skor: {evaluation.total_score}/{evaluation.max_score}
                      </Typography>
                      <Typography variant="body2">
                        ({Math.round((evaluation.total_score / evaluation.max_score) * 100)}%)
                      </Typography>
                    </Box>

                    {evaluation.teacher_feedback && (
                      <Box sx={{ mt: 1, p: 1, bgcolor: 'grey.50', borderRadius: 1 }}>
                        <Typography variant="caption" color="text.secondary" gutterBottom>
                          Feedback:
                        </Typography>
                        <Typography variant="body2" sx={{ fontSize: '0.875rem' }}>
                          {evaluation.teacher_feedback}
                        </Typography>
                      </Box>
                    )}
                  </Box>
                </TimelineContent>
              </TimelineItem>
            ))}
          </Timeline>
        )}
      </CardContent>
    </Card>
  );
};

export default RevisionHistory;
