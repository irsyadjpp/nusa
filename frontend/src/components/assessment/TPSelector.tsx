/**
 * TP Selector Component
 * Dropdown/search for selecting Teaching Plan (TP)
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  TextField,
  Autocomplete,
  Typography,
  Chip,
  Card,
  CardContent,
  IconButton,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
} from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import VisibilityIcon from '@mui/icons-material/Visibility';
import { getTPs } from '@/api/tp';
import { TP } from '@/api/tp';

interface TPSelectorProps {
  value?: string;
  onChange: (tpId: string, tp: TP) => void;
  disabled?: boolean;
  error?: boolean;
  helperText?: string;
  subjectId?: string;
  phaseId?: string;
}

const TPSelector: React.FC<TPSelectorProps> = ({
  value,
  onChange,
  disabled = false,
  error = false,
  helperText,
  subjectId,
  phaseId,
}) => {
  const [tps, setTps] = useState<TP[]>([]);
  const [loading, setLoading] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedTP, setSelectedTP] = useState<TP | null>(null);
  const [previewOpen, setPreviewOpen] = useState(false);

  useEffect(() => {
    loadTPs();
  }, [subjectId, phaseId]);

  const loadTPs = async () => {
    setLoading(true);
    try {
      const params: any = {};
      if (subjectId) params.subject_id = subjectId;
      if (phaseId) params.phase_id = phaseId;
      const data = await getTPs(params);
      setTps(data);
    } catch (error) {
      console.error('Error loading TPs:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleTPChange = (_: any, newValue: TP | null) => {
    if (newValue) {
      setSelectedTP(newValue);
      onChange(newValue.id, newValue);
    } else {
      setSelectedTP(null);
      onChange('', {} as TP);
    }
  };

  const handlePreview = () => {
    if (selectedTP) {
      setPreviewOpen(true);
    }
  };

  const filteredTPs = tps.filter((tp) =>
    tp.title.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <Box>
      <Autocomplete
        value={selectedTP}
        onChange={handleTPChange}
        options={filteredTPs}
        getOptionLabel={(option) => option.title}
        renderInput={(params) => (
          <TextField
            {...params}
            label="Pilih Rencana Pelaksanaan Pembelajaran (TP)"
            placeholder="Cari TP..."
            error={error}
            helperText={helperText}
            disabled={disabled || loading}
            InputProps={{
              ...params.InputProps,
              startAdornment: <SearchIcon sx={{ mr: 1, color: 'text.secondary' }} />,
            }}
          />
        )}
        renderOption={(props, option) => (
          <Box component="li" {...props}>
            <Box sx={{ flexGrow: 1 }}>
              <Typography variant="body1" fontWeight="medium">
                {option.title}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {option.sequence_number}. TP - {option.status}
              </Typography>
            </Box>
          </Box>
        )}
        loading={loading}
        disabled={disabled}
        fullWidth
      />

      {selectedTP && (
        <Box sx={{ mt: 2, display: 'flex', alignItems: 'center', gap: 2 }}>
          <Chip
            label={`TP ${selectedTP.sequence_number}`}
            size="small"
            color="primary"
          />
          <Chip
            label={selectedTP.status}
            size="small"
            color={selectedTP.status === 'APPROVED' ? 'success' : 'default'}
          />
          <IconButton onClick={handlePreview} size="small">
            <VisibilityIcon />
          </IconButton>
        </Box>
      )}

      {/* Preview Dialog */}
      <Dialog
        open={previewOpen}
        onClose={() => setPreviewOpen(false)}
        maxWidth="md"
        fullWidth
      >
        <DialogTitle>Preview TP</DialogTitle>
        <DialogContent>
          {selectedTP && (
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  {selectedTP.title}
                </Typography>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  Urutan: {selectedTP.sequence_number}
                </Typography>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  Status: {selectedTP.status}
                </Typography>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  Estimasi Minggu: {selectedTP.estimated_weeks}
                </Typography>
                {selectedTP.success_criteria && (
                  <Box sx={{ mt: 2 }}>
                    <Typography variant="subtitle2" gutterBottom>
                      Kriteria Ketuntasan:
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      {JSON.stringify(selectedTP.success_criteria, null, 2)}
                    </Typography>
                  </Box>
                )}
              </CardContent>
            </Card>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setPreviewOpen(false)}>Tutup</Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default TPSelector;
