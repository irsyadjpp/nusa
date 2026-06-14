/**
 * Evidence List Page
 * List view for Evidence
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
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
import { getEvidences } from '@/api/evidence';
import { Evidence, EvidenceType } from '@/shared/types/domain';
// import EvidenceReview from '@/components/evidence/EvidenceReview'; // TODO: Implement

const EvidenceListPage: React.FC = () => {
  const navigate = useNavigate();
  const [evidences, setEvidences] = useState<Evidence[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedStudent, setSelectedStudent] = useState<string>('');
  const [selectedAssessment, setSelectedAssessment] = useState<string>('');
  const [selectedType, setSelectedType] = useState<EvidenceType | ''>('');
  const [selectedStatus, setSelectedStatus] = useState<string>('');

  useEffect(() => {
    loadEvidences();
  }, [selectedStudent, selectedAssessment, selectedType, selectedStatus]);

  const loadEvidences = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await getEvidences({
        student_id: selectedStudent || undefined,
        assessment_id: selectedAssessment || undefined,
        evidence_type: selectedType || undefined,
        status: selectedStatus || undefined,
      });
      setEvidences(data);
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data Bukti');
    } finally {
      setLoading(false);
    }
  };

  const filteredEvidences = evidences.filter((evidence) =>
    evidence.student_name?.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4">Bukti Pembelajaran</Typography>
        <Button
          variant="contained"
          startIcon={<Add />}
          onClick={() => navigate('/evidence/upload')}
        >
          Upload Bukti Baru
        </Button>
      </Box>

      <Card sx={{ mb: 3 }}>
        <CardContent>
          <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2 }}>
            <Box sx={{ width: { xs: '100%', sm: '33.33%' } }}>
              <TextField
                fullWidth
                label="Cari Bukti"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                InputProps={{
                  startAdornment: <Search sx={{ mr: 1, color: 'text.secondary' }} />,
                }}
              />
            </Box>
            <Box sx={{ width: { xs: '100%', sm: '16.67%' } }}>
              <FormControl fullWidth>
                <InputLabel>Tipe</InputLabel>
                <Select
                  value={selectedType}
                  label="Tipe"
                  onChange={(e) => setSelectedType(e.target.value)}
                >
                  <MenuItem value="">Semua</MenuItem>
                  <MenuItem value="STUDENT_WORK">Karya Siswa</MenuItem>
                  <MenuItem value="ASSESSMENT_RESULT">Hasil Asesmen</MenuItem>
                  <MenuItem value="OBSERVATION">Observasi</MenuItem>
                </Select>
              </FormControl>
            </Box>
            <Box sx={{ width: { xs: '100%', sm: '16.67%' } }}>
              <FormControl fullWidth>
                <InputLabel>Status</InputLabel>
                <Select
                  value={selectedStatus}
                  label="Status"
                  onChange={(e) => setSelectedStatus(e.target.value)}
                >
                  <MenuItem value="">Semua</MenuItem>
                  <MenuItem value="SUBMITTED">Dikirim</MenuItem>
                  <MenuItem value="PENDING">Pending</MenuItem>
                  <MenuItem value="EVALUATED">Dievaluasi</MenuItem>
                  <MenuItem value="REJECTED">Ditolak</MenuItem>
                </Select>
              </FormControl>
            </Box>
            <Box sx={{ width: { xs: '100%', sm: '33.33%' } }}>
              <Button
                fullWidth
                variant="outlined"
                startIcon={<FilterList />}
                onClick={() => {
                  setSearchTerm('');
                  setSelectedStudent('');
                  setSelectedAssessment('');
                  setSelectedType('');
                  setSelectedStatus('');
                }}
              >
                Reset Filter
              </Button>
            </Box>
          </Box>
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
        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
          {filteredEvidences.map((evidence) => (
            <Box sx={{ width: { xs: '100%', sm: '50%', md: '33.33%' } }} key={evidence.id}>
              <Card>
                <CardContent>
                  <Typography variant="subtitle1" gutterBottom>
                    Evidence: {evidence.id.substring(0, 8)}...
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Student: {evidence.student_id.substring(0, 8)}...
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Status: {evidence.status}
                  </Typography>
                  <Box sx={{ mt: 2, display: 'flex', gap: 1 }}>
                    <Button size="small" onClick={() => navigate(`/evidence/${evidence.id}`)}>
                      View
                    </Button>
                  </Box>
                </CardContent>
              </Card>
            </Box>
          ))}
        </Box>
      )}

      {!loading && filteredEvidences.length === 0 && (
        <Alert severity="info">
          Tidak ada bukti yang ditemukan
        </Alert>
      )}
    </Box>
  );
};

export default EvidenceListPage;
