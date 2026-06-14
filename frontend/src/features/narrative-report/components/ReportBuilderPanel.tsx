/**
 * Report Builder Panel Component
 * Panel for building narrative reports
 */

import { Box, Typography, Paper, Divider, Button } from '@mui/material';
import { NarrativeReport } from '@/shared/types/domain';

interface ReportBuilderPanelProps {
  report: NarrativeReport;
  onEdit?: () => void;
}

export const ReportBuilderPanel = ({ report, onEdit }: ReportBuilderPanelProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        {report.title || `Report ${report.id}`}
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Status
          </Typography>
          <Typography variant="body1">{report.status}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Student ID
          </Typography>
          <Typography variant="body1">{report.student_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Period
          </Typography>
          <Typography variant="body1">{report.period || report.period_id}</Typography>
        </Box>

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Content
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {typeof report.content === 'string'
              ? report.content
              : typeof report.narrative_content === 'string'
              ? report.narrative_content
              : report.content || report.narrative_content
              ? JSON.stringify(report.content || report.narrative_content, null, 2)
              : 'No content yet'}
          </Typography>
        </Box>

        {onEdit && (
          <Button variant="outlined" onClick={onEdit}>
            Edit Report
          </Button>
        )}
      </Box>
    </Paper>
  );
};
