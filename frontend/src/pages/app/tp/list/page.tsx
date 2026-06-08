/**
 * TP List Page
 * List view for Teaching Plans (TP)
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Grid,
  Card,
  CardContent,
  Chip,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  Add,
  Search,
  FilterList,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { getTPs, getTPSets } from '@/api/tp';
import { TP, TPSet } from '@/api/tp';

const TPListPage: React.FC = () => {
  const navigate = useNavigate();
  const [tps, setTps] = useState<TP[]>([]);
  const [tpSets, setTPSets] = useState<TPSet[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedTPSet, setSelectedTPSet] = useState<string>('');
  const [selectedStatus, setSelectedStatus] = useState<string>('');
  const [selectedSubject, setSelectedSubject] = useState<string>('');

  useEffect(() => {
    loadData();
  }, [selectedTPSet, selectedStatus, selectedSubject]);

  const loadData = async () => {
    setLoading(true);
    setError(null);
    try {
      const [tpsData, tpSetsData] = await Promise.all([
        getTPs({
          tp_set_id: selectedTPSet || undefined,
          status: selectedStatus || undefined,
          subject_id: selectedSubject || undefined,
        }),
        getTPSets(),
      ]);
      setTps(tpsData);
      setTPSets(tpSetsData);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data TP');
    } finally {
      setLoading(false);
    }
  };

  const filteredTPs = tps.filter((tp) =>
    tp.title.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const getStatusColor = (status: string): 'success' | 'warning' | 'error' | 'info' | 'default' => {
    switch (status) {
      case 'APPROVED':
        return 'success';
      case 'PENDING':
        return 'warning';
      case 'REJECTED':
        return 'error';
      case 'DRAFT':
        return 'info';
      default:
        return 'default';
    }
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4">Rencana Pelaksanaan Pembelajaran (TP)</Typography>
        <Button
          variant="contained"
          startIcon={<Add />}
          onClick={() => navigate('/tp/create')}
        >
          Buat TP Baru
        </Button>
      </Box>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Grid container spacing={2}>
            <Grid item xs={12} sm={4}>
              <TextField
                fullWidth
                label="Cari TP"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                InputProps={{
                  startAdornment: <Search sx={{ mr: 1, color: 'text.secondary' }} />,
                }}
              />
            </Grid>
            <Grid item xs={12} sm={3}>
              <FormControl fullWidth>
                <InputLabel>Set TP</InputLabel>
                <Select
                  value={selectedTPSet}
                  label="Set TP"
                  onChange={(e) => setSelectedTPSet(e.target.value)}
                >
                  <MenuItem value="">Semua</MenuItem>
                  {tpSets.map((tpSet) => (
                    <MenuItem key={tpSet.id} value={tpSet.id}>
                      {tpSet.name}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} sm={2}>
              <FormControl fullWidth>
                <InputLabel>Status</InputLabel>
                <Select
                  value={selectedStatus}
                  label="Status"
                  onChange={(e) => setSelectedStatus(e.target.value)}
                >
                  <MenuItem value="">Semua</MenuItem>
                  <MenuItem value="DRAFT">Draft</MenuItem>
                  <MenuItem value="PENDING">Pending</MenuItem>
                  <MenuItem value="APPROVED">Disetujui</MenuItem>
                  <MenuItem value="REJECTED">Ditolak</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} sm={3}>
              <Button
                fullWidth
                variant="outlined"
                startIcon={<FilterList />}
                onClick={() => {
                  setSearchTerm('');
                  setSelectedTPSet('');
                  setSelectedStatus('');
                  setSelectedSubject('');
                }}
              >
                Reset Filter
              </Button>
            </Grid>
          </Grid>
        </CardContent>
      </Card>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      {loading ? (
        <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
          <CircularProgress />
        </Box>
      ) : (
        <Grid container spacing={3}>
          {filteredTPs.map((tp) => (
            <Grid item xs={12} sm={6} md={4} key={tp.id}>
              <Card
                sx={{ cursor: 'pointer', height: '100%' }}
                onClick={() => navigate(`/tp/${tp.id}`)}
              >
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    {tp.title}
                  </Typography>
                  <Typography variant="body2" color="text.secondary" gutterBottom>
                    TP #{tp.sequence_number}
                  </Typography>
                  <Box sx={{ display: 'flex', gap: 1, mb: 2 }}>
                    <Chip
                      label={tp.status}
                      color={getStatusColor(tp.status)}
                      size="small"
                    />
                    {tp.success_criteria && (
                      <Chip
                        label="KKTP"
                        color="primary"
                        size="small"
                        variant="outlined"
                      />
                    )}
                  </Box>
                  <Typography variant="caption" color="text.secondary">
                    Estimasi: {tp.estimated_weeks} minggu
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      )}

      {!loading && filteredTPs.length === 0 && (
        <Alert severity="info">
          Tidak ada TP yang ditemukan
        </Alert>
      )}
    </Box>
  );
};

export default TPListPage;
