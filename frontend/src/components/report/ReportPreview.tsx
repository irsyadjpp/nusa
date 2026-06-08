/**
 * Report Preview Component
 * Preview of final report
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Button,
  Divider,
  Paper,
} from '@mui/material';
import { Print } from '@mui/icons-material';
import ReactQuill from 'react-quill-new';
import 'react-quill-new/dist/quill.bubble.css';

interface ReportPreviewProps {
  studentName: string;
  periodStart: string;
  periodEnd: string;
  narrativeContent: string;
  achievementSummary?: any;
  onPrint?: () => void;
}

const ReportPreview: React.FC<ReportPreviewProps> = ({
  studentName,
  periodStart,
  periodEnd,
  narrativeContent,
  achievementSummary,
  onPrint,
}) => {
  return (
    <Card>
      <CardContent>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
          <Typography variant="h6">Preview Rapor</Typography>
          {onPrint && (
            <Button
              variant="contained"
              startIcon={<Print />}
              onClick={onPrint}
            >
              Cetak
            </Button>
          )}
        </Box>

        <Paper
          sx={{
            p: 4,
            bgcolor: 'white',
            minHeight: '600px',
            maxHeight: '800px',
            overflowY: 'auto',
          }}
          id="report-preview"
        >
          <Typography variant="h4" align="center" gutterBottom>
            Rapor Naratif
          </Typography>
          <Typography variant="h5" align="center" gutterBottom>
            {studentName}
          </Typography>
          <Typography variant="body1" align="center" color="text.secondary" gutterBottom>
            Periode: {new Date(periodStart).toLocaleDateString('id-ID')} -{' '}
            {new Date(periodEnd).toLocaleDateString('id-ID')}
          </Typography>

          <Divider sx={{ my: 3 }} />

          {achievementSummary && (
            <Box sx={{ mb: 3 }}>
              <Typography variant="h6" gutterBottom>
                Ringkasan Pencapaian
              </Typography>
              <Box sx={{ mb: 2 }}>
                <Typography variant="body2" gutterBottom>
                  Penguasaan Keseluruhan: {achievementSummary.overall_mastery.toFixed(1)}%
                </Typography>
              </Box>
            </Box>
          )}

          <Divider sx={{ my: 3 }} />

          <Typography variant="h6" gutterBottom>
            Narasi Guru
          </Typography>
          <ReactQuill
            theme="bubble"
            value={narrativeContent}
            readOnly
            style={{ minHeight: '300px' }}
          />

          <Divider sx={{ my: 3 }} />

          <Box sx={{ mt: 4, pt: 2, borderTop: '1px solid #e0e0e0' }}>
            <Typography variant="body2" color="text.secondary" align="center">
              Rapor ini dibuat secara otomatis oleh NUSA Education Operating System
            </Typography>
            <Typography variant="caption" color="text.secondary" align="center" display="block">
              {new Date().toLocaleDateString('id-ID', {
                weekday: 'long',
                year: 'numeric',
                month: 'long',
                day: 'numeric',
              })}
            </Typography>
          </Box>
        </Paper>
      </CardContent>
    </Card>
  );
};

export default ReportPreview;
