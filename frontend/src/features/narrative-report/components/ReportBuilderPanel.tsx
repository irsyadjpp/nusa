/**
 * Report Builder Panel Component
 * Panel for building narrative reports
 */

import { Box, Typography, Paper, Divider, Stack } from '@mui/material';
import { NarrativeReport } from '@/api/narrative-report';

interface ReportBuilderPanelProps {
  report: NarrativeReport;
  onEdit?: () => void;
}

export const ReportBuilderPanel = ({ report, onEdit }: ReportBuilderPanelProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        {report.title}
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={2}>
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
          <Typography variant="body1">{report.period}</Typography>
        </Box>

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Content
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {report.content || 'No content yet'}
          </Typography>
        </Box>

        {onEdit && (
          <Button variant="outlined" onClick={onEdit}>
            Edit Report
          </Button>
        )}
      </Stack>
    </Paper>
  );
};
