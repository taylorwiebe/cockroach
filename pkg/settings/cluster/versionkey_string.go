// Code generated by "stringer -type=VersionKey"; DO NOT EDIT.

package cluster

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Version2_0-0]
	_ = x[VersionImportSkipRecords-1]
	_ = x[VersionProposedTSLeaseRequest-2]
	_ = x[VersionImportFormats-3]
	_ = x[Version2_1-4]
	_ = x[VersionCascadingZoneConfigs-5]
	_ = x[VersionLoadSplits-6]
	_ = x[VersionExportStorageWorkload-7]
	_ = x[VersionLazyTxnRecord-8]
	_ = x[VersionSequencedReads-9]
	_ = x[VersionUnreplicatedRaftTruncatedState-10]
	_ = x[VersionCreateStats-11]
	_ = x[VersionDirectImport-12]
	_ = x[VersionSideloadedStorageNoReplicaID-13]
	_ = x[VersionPushTxnToInclusive-14]
	_ = x[VersionSnapshotsWithoutLog-15]
	_ = x[Version19_1-16]
	_ = x[VersionStart19_2-17]
	_ = x[VersionQueryTxnTimestamp-18]
	_ = x[VersionStickyBit-19]
	_ = x[VersionParallelCommits-20]
}

const _VersionKey_name = "Version2_0VersionImportSkipRecordsVersionProposedTSLeaseRequestVersionImportFormatsVersion2_1VersionCascadingZoneConfigsVersionLoadSplitsVersionExportStorageWorkloadVersionLazyTxnRecordVersionSequencedReadsVersionUnreplicatedRaftTruncatedStateVersionCreateStatsVersionDirectImportVersionSideloadedStorageNoReplicaIDVersionPushTxnToInclusiveVersionSnapshotsWithoutLogVersion19_1VersionStart19_2VersionQueryTxnTimestampVersionStickyBitVersionParallelCommits"

var _VersionKey_index = [...]uint16{0, 10, 34, 63, 83, 93, 120, 137, 165, 185, 206, 243, 261, 280, 315, 340, 366, 377, 393, 417, 433, 455}

func (i VersionKey) String() string {
	if i < 0 || i >= VersionKey(len(_VersionKey_index)-1) {
		return "VersionKey(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VersionKey_name[_VersionKey_index[i]:_VersionKey_index[i+1]]
}
