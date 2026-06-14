/**
 * Report Detail Page
 * Detail view for Narrative Report
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  CircularProgress,
  Alert,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Card,
  CardContent,
} from '@mui/material';
import {
  ArrowBack,
  Edit,
  Delete,
  Print,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import NarrativeReportEditor from '@/components/report/NarrativeReportEditor';
import ReportPreview from '@/components/report/ReportPreview';
import ReportPublish from '@/components/report/ReportPublish';

const ReportDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [report, setReport] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [deleting, setDeleting] = useState(false);
  const [editing, setEditing] = useState(false);
  const [narrativeContent, setNarrativeContent] = useState('');

  useEffect(() => {
    if (id) {
      loadReport(id);
    }
  }, [id]);

  const loadReport = async (reportId: string) => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch(`/api/reports/${reportId}`);
      const data = await response.json();
      setReport(data);
      setNarrativeContent(data.narrative_content || '');
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data Rapor');
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async () => {
    if (!id) return;
    setDeleting(true);
    try {
      await fetch(`/api/reports/${id}`, { method: 'DELETE' });
      navigate('/reports');
    } catch (err: any) {
      setError(err.message || 'Gagal menghapus rapor');
      setDeleteDialogOpen(false);
    } finally {
      setDeleting(false);
    }
  };

  const handleSave = async () => {
    if (!id) return;
    try {
      await fetch(`/api/reports/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ narrative_content: narrativeContent }),
      });
      loadReport(id);
      setEditing(false);
    } catch (err: any) {
      setError(err.message || 'Gagal menyimpan rapor');
    }
  };

  const handlePublish = async (reportId: string, publishData: any) => {
    try {
      await fetch(`/api/reports/${reportId}/publish`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(publishData),
      });
      loadReport(id!);
    } catch (err: any) {
      setError(err.message || 'Gagal mempublikasikan rapor');
    }
  };

  const handlePrint = () => {
    window.print();
  };

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error || !report) {
    return (
      <Alert severity="error">
        {error || 'Rapor tidak ditemukan'}
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
            onClick={() => navigate('/reports')}
          >
            Kembali
          </Button>
          <Typography variant="h4">Detail Rapor</Typography>
        </Box>
        <Box sx={{ display: 'flex', gap: 1 }}>
          {report.status === 'DRAFT' && (
            <>
              <Button
                variant="outlined"
                startIcon={<Edit />}
                onClick={() => setEditing(true)}
              >
                Edit
              </Button>
              <ReportPublish
                reportId={report.id}
                studentName={report.student_name}
                onPublish={handlePublish}
              />
            </>
          )}
          <Button
            variant="outlined"
            startIcon={<Print />}
            onClick={handlePrint}
          >
            Cetak
          </Button>
          <Button
            variant="outlined"
            color="error"
            startIcon={<Delete />}
            onClick={() => setDeleteDialogOpen(true)}
          >
            Hapus
          </Button>
        </Box>
      </Box>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      {!editing ? (
        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
          <Box sx={{ width: { xs: '100%', md: '66.67%' }}}>
            <Card sx={{ mb: 3 }}>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  Narasi Guru
                </Typography>
                <Box
                  sx={{
                    p: 2,
                    border: 1,
                    borderColor: 'divider',
                    borderRadius: 1,
                    minHeight: '200px',
                    whiteSpace: 'pre-wrap',
                  }}
                >
                  {narrativeContent || 'Belum ada narasi'}
                </Box>
              </CardContent>
            </Card>
          </Box>

          <Box sx={{ width: { xs: '100%', md: '33.33%' }}}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  Informasi
                </Typography>
                <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                  <Box>
                    <Typography variant="caption" color="text.secondary">
                      Siswa
                    </Typography>
                    <Typography variant="body2">
                      {report.student_name}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="caption" color="text.secondary">
                      Periode
                    </Typography>
                    <Typography variant="body2">
                      {new Date(report.period_start).toLocaleDateString('id-ID')} -{' '}
                      {new Date(report.period_end).toLocaleDateString('id-ID')}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="caption" color="text.secondary">
                      Status
                    </Typography>
                    <Typography variant="body2">
                      {report.status}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="caption" color="text.secondary">
                      Dibuat
                    </Typography>
                    <Typography variant="body2">
                      {new Date(report.created_at).toLocaleString('id-ID')}
                    </Typography>
                  </Box>
                  {report.published_at && (
                    <Box>
                      <Typography variant="caption" color="text.secondary">
                        Dipublikasikan
                      </Typography>
                      <Typography variant="body2">
                        {new Date(report.published_at).toLocaleString('id-ID')}
                      </Typography>
                    </Box>
                  )}
                </Box>
              </CardContent>
            </Card>
          </Box>
        </Box>
      ) : editing ? (
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>
              Edit Narasi Guru
            </Typography>
            <NarrativeReportEditor
              value={narrativeContent}
              onChange={setNarrativeContent}
              onSave={handleSave}
              onCancel={() => setEditing(false)}
            />
          </CardContent>
        </Card>
      ) : (
        <ReportPreview
          studentName={report.student_name}
          periodStart={report.period_start}
          periodEnd={report.period_end}
          narrativeContent={narrativeContent}
          onPrint={handlePrint}
        />
      )}

      <Dialog open={deleteDialogOpen} onClose={() => setDeleteDialogOpen(false)}>
        <DialogTitle>Hapus Rapor</DialogTitle>
        <DialogContent>
          Apakah Anda yakin ingin menghapus rapor ini? Tindakan ini tidak dapat dibatalkan.
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setDeleteDialogOpen(false)}>Batal</Button>
          <Button
            variant="contained"
            color="error"
            onClick={handleDelete}
            disabled={deleting}
          >
            {deleting ? 'Menghapus...' : 'Hapus'}
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default ReportDetailPage;
