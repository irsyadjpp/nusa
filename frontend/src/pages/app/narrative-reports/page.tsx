import React, { useState } from 'react';
import {
  Box,
  Typography,
  Button,
  Paper,
  TextField,
  CircularProgress,
  Alert,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Chip,
  IconButton,
  Container,
  Grid,
} from '@mui/material';
import { useNavigate } from 'react-router-dom';
import DeleteIcon from '@mui/icons-material/Delete';
import VisibilityIcon from '@mui/icons-material/Visibility';
import EditIcon from '@mui/icons-material/Edit';
import { useNarrativeReports } from '@/services/queries/NarrativeReportQueryService';
import {
  deleteNarrativeReport,
  publishNarrativeReport,
} from '@/api/narrative-report';

const NarrativeReportsPage: React.FC = () => {
  const navigate = useNavigate();
  const [filters, setFilters] = useState({
    student_id: '',
    period_id: '',
    status: '',
  });

  const {
    data: reports,
    isLoading,
    error,
    refetch,
  } = useNarrativeReports(filters);

  const handleFilterChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFilters({
      ...filters,
      [e.target.name]: e.target.value,
    });
  };

  const handleApplyFilters = async () => {
    await refetch();
  };

  const handleDelete = async (id: string) => {
    if (window.confirm('Are you sure you want to delete this narrative report?')) {
      try {
        await deleteNarrativeReport(id);
        refetch();
      } catch (error) {
        console.error('Failed to delete narrative report:', error);
      }
    }
  };

  const handlePublish = async (id: string) => {
    try {
      await publishNarrativeReport(id);
      refetch();
    } catch (error) {
      console.error('Failed to publish narrative report:', error);
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'PUBLISHED':
        return 'success';
      case 'DRAFT':
        return 'default';
      case 'PENDING_REVIEW':
        return 'warning';
      default:
        return 'default';
    }
  };

  if (isLoading) {
    return (
      <Container maxWidth="xl" sx={{ py: 4 }}>
        <Box sx={{ display: 'flex', justifyContent: 'center', py: 8 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (error) {
    return (
      <Container maxWidth="xl" sx={{ py: 4 }}>
        <Alert severity="error">{error.message || 'Failed to load narrative reports'}</Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="xl" sx={{ py: 4 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 4 }}>
        <Typography variant="h4" component="h1">
          Narrative Reports
        </Typography>
        <Button
          variant="contained"
          onClick={() => navigate('/narrative-reports/generate')}
        >
          Generate Report
        </Button>
      </Box>

      <Paper sx={{ p: 3, mb: 3 }}>
        <Grid container spacing={2} alignItems="center">
          <Grid size={{ xs: 12, md: 3 }}>
            <TextField
              label="Student ID"
              name="student_id"
              value={filters.student_id}
              onChange={handleFilterChange}
              fullWidth
              size="small"
            />
          </Grid>
          <Grid size={{ xs: 12, md: 3 }}>
            <TextField
              label="Period ID"
              name="period_id"
              value={filters.period_id}
              onChange={handleFilterChange}
              fullWidth
              size="small"
            />
          </Grid>
          <Grid size={{ xs: 12, md: 3 }}>
            <TextField
              label="Status"
              name="status"
              value={filters.status}
              onChange={handleFilterChange}
              fullWidth
              size="small"
            />
          </Grid>
          <Grid size={{ xs: 12, md: 3 }}>
            <Button variant="outlined" onClick={handleApplyFilters}>
              Apply Filters
            </Button>
          </Grid>
        </Grid>
      </Paper>

      {reports && reports.length > 0 ? (
        <Paper>
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>ID</TableCell>
                  <TableCell>Student ID</TableCell>
                  <TableCell>Period ID</TableCell>
                  <TableCell>Status</TableCell>
                  <TableCell>Created By</TableCell>
                  <TableCell>Published At</TableCell>
                  <TableCell>Created At</TableCell>
                  <TableCell>Actions</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {reports.map((report: any) => (
                  <TableRow key={report.id}>
                    <TableCell>{report.id}</TableCell>
                    <TableCell>{report.student_id}</TableCell>
                    <TableCell>{report.period_id}</TableCell>
                    <TableCell>
                      <Chip
                        label={report.status}
                        color={getStatusColor(report.status) as any}
                        size="small"
                      />
                    </TableCell>
                    <TableCell>{report.created_by}</TableCell>
                    <TableCell>
                      {report.published_at
                        ? new Date(report.published_at).toLocaleDateString()
                        : '-'}
                    </TableCell>
                    <TableCell>
                      {new Date(report.created_at).toLocaleDateString()}
                    </TableCell>
                    <TableCell>
                      <IconButton
                        size="small"
                        onClick={() => navigate(`/narrative-reports/${report.id}`)}
                        title="View"
                      >
                        <VisibilityIcon />
                      </IconButton>
                      <IconButton
                        size="small"
                        onClick={() => navigate(`/narrative-reports/${report.id}/edit`)}
                        title="Edit"
                      >
                        <EditIcon />
                      </IconButton>
                      {report.status === 'DRAFT' && (
                        <Button
                          size="small"
                          onClick={() => handlePublish(report.id)}
                          variant="text"
                        >
                          Publish
                        </Button>
                      )}
                      <IconButton
                        size="small"
                        onClick={() => handleDelete(report.id)}
                        title="Delete"
                      >
                        <DeleteIcon />
                      </IconButton>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Paper>
      ) : (
        <Alert severity="info">No narrative reports found</Alert>
      )}
    </Container>
  );
};

export default NarrativeReportsPage;
