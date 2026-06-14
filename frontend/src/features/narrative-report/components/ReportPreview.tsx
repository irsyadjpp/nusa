/**
 * Report Preview Component
 * Displays a preview of the narrative report
 */

import { Box, Typography, Paper, Divider, Button, Stack } from '@mui/material';
import { Print as PrintIcon, Download as DownloadIcon } from '@mui/icons-material';
import { NarrativeReport } from '@/shared/types/domain';

interface ReportPreviewProps {
  report: NarrativeReport;
  onPrint?: () => void;
  onDownload?: () => void;
}

export const ReportPreview = ({ report, onPrint, onDownload }: ReportPreviewProps) => {
  return (
    <Paper elevation={2} sx={{ p: 4 }}>
      <Stack direction="row" justifyContent="space-between" alignItems="center" mb={3}>
        <Typography variant="h5">Report Preview</Typography>
        <Stack direction="row" spacing={2}>
          {onPrint && (
            <Button variant="outlined" startIcon={<PrintIcon />} onClick={onPrint}>
              Print
            </Button>
          )}
          {onDownload && (
            <Button variant="contained" startIcon={<DownloadIcon />} onClick={onDownload}>
              Download
            </Button>
          )}
        </Stack>
      </Stack>

      <Divider sx={{ mb: 3 }} />

      <Box sx={{ fontFamily: 'serif', lineHeight: 1.8 }}>
        <Typography variant="h4" gutterBottom>
          {report.title || `Narrative Report`}
        </Typography>

        <Typography variant="subtitle1" color="textSecondary" gutterBottom>
          Student ID: {report.student_id}
        </Typography>

        <Typography variant="subtitle1" color="textSecondary" gutterBottom>
          Period: {report.period || report.period_id}
        </Typography>

        <Divider sx={{ my: 3 }} />

        <Typography variant="body1" whiteSpace="pre-wrap">
          {typeof report.content === 'string'
            ? report.content
            : typeof report.narrative_content === 'string'
            ? report.narrative_content
            : report.content || report.narrative_content
            ? JSON.stringify(report.content || report.narrative_content, null, 2)
            : 'No content available'}
        </Typography>
      </Box>

      <Divider sx={{ my: 3 }} />

      <Typography variant="caption" color="textSecondary">
        Generated on {new Date(report.created_at).toLocaleString()}
      </Typography>
    </Paper>
  );
};
