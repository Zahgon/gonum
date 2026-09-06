package lapack

import "gonum.org/v1/gonum/blas"

type Complex128 interface{}

type Float64 interface {
	Dgecon(norm MatrixNorm, n int, a []float64, lda int, anorm float64, work []float64, iwork []int) float64
	Dgeev(jobvl LeftEVJob, jobvr RightEVJob, n int, a []float64, lda int, wr, wi []float64, vl []float64, ldvl int, vr []float64, ldvr int, work []float64, lwork int) (first int)
	Dgels(trans blas.Transpose, m, n, nrhs int, a []float64, lda int, b []float64, ldb int, work []float64, lwork int) bool
	Dgelqf(m, n int, a []float64, lda int, tau, work []float64, lwork int)
	Dgeqp3(m, n int, a []float64, lda int, jpvt []int, tau, work []float64, lwork int)
	Dgeqrf(m, n int, a []float64, lda int, tau, work []float64, lwork int)
	Dgesvd(jobU, jobVT SVDJob, m, n int, a []float64, lda int, s, u []float64, ldu int, vt []float64, ldvt int, work []float64, lwork int) (ok bool)
	Dgetrf(m, n int, a []float64, lda int, ipiv []int) (ok bool)
	Dgetri(n int, a []float64, lda int, ipiv []int, work []float64, lwork int) (ok bool)
	Dgetrs(trans blas.Transpose, n, nrhs int, a []float64, lda int, ipiv []int, b []float64, ldb int)
	Dggsvd3(jobU, jobV, jobQ GSVDJob, m, n, p int, a []float64, lda int, b []float64, ldb int, alpha, beta, u []float64, ldu int, v []float64, ldv int, q []float64, ldq int, work []float64, lwork int, iwork []int) (k, l int, ok bool)
	Dlantr(norm MatrixNorm, uplo blas.Uplo, diag blas.Diag, m, n int, a []float64, lda int, work []float64) float64
	Dlange(norm MatrixNorm, m, n int, a []float64, lda int, work []float64) float64
	Dlansy(norm MatrixNorm, uplo blas.Uplo, n int, a []float64, lda int, work []float64) float64
	Dlapmr(forward bool, m, n int, x []float64, ldx int, k []int)
	Dlapmt(forward bool, m, n int, x []float64, ldx int, k []int)
	Dorgqr(m, n, k int, a []float64, lda int, tau, work []float64, lwork int)
	Dormqr(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int)
	Dorglq(m, n, k int, a []float64, lda int, tau, work []float64, lwork int)
	Dormlq(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int)
	Dpbcon(uplo blas.Uplo, n, kd int, ab []float64, ldab int, anorm float64, work []float64, iwork []int) float64
	Dpbtrf(uplo blas.Uplo, n, kd int, ab []float64, ldab int) (ok bool)
	Dpbtrs(uplo blas.Uplo, n, kd, nrhs int, ab []float64, ldab int, b []float64, ldb int)
	Dpocon(uplo blas.Uplo, n int, a []float64, lda int, anorm float64, work []float64, iwork []int) float64
	Dpotrf(ul blas.Uplo, n int, a []float64, lda int) (ok bool)
	Dpotri(ul blas.Uplo, n int, a []float64, lda int) (ok bool)
	Dpotrs(ul blas.Uplo, n, nrhs int, a []float64, lda int, b []float64, ldb int)
	Dpstrf(uplo blas.Uplo, n int, a []float64, lda int, piv []int, tol float64, work []float64) (rank int, ok bool)
	Dsyev(jobz EVJob, uplo blas.Uplo, n int, a []float64, lda int, w, work []float64, lwork int) (ok bool)
	Dtbtrs(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, kd, nrhs int, a []float64, lda int, b []float64, ldb int) (ok bool)
	Dtrcon(norm MatrixNorm, uplo blas.Uplo, diag blas.Diag, n int, a []float64, lda int, work []float64, iwork []int) float64
	Dtrtri(uplo blas.Uplo, diag blas.Diag, n int, a []float64, lda int) (ok bool)
	Dtrtrs(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, nrhs int, a []float64, lda int, b []float64, ldb int) (ok bool)
}

type Direct byte

const (
	Forward  Direct = 'F'
	Backward Direct = 'B'
)

type Sort byte

const (
	SortIncreasing Sort = 'I'
	SortDecreasing Sort = 'D'
)

type StoreV byte

const (
	ColumnWise StoreV = 'C'
	RowWise    StoreV = 'R'
)

type MatrixNorm byte

const (
	MaxAbs       MatrixNorm = 'M'
	MaxColumnSum MatrixNorm = 'O'
	MaxRowSum    MatrixNorm = 'I'
	Frobenius    MatrixNorm = 'F'
)

type MatrixType byte

const (
	General  MatrixType = 'G'
	UpperTri MatrixType = 'U'
	LowerTri MatrixType = 'L'
)

type Pivot byte

const (
	Variable Pivot = 'V'
	Top      Pivot = 'T'
	Bottom   Pivot = 'B'
)

type ApplyOrtho byte

const (
	ApplyP ApplyOrtho = 'P'
	ApplyQ ApplyOrtho = 'Q'
)

type GenOrtho byte

const (
	GeneratePT GenOrtho = 'P'
	GenerateQ  GenOrtho = 'Q'
)

type SVDJob byte

const (
	SVDAll       SVDJob = 'A'
	SVDStore     SVDJob = 'S'
	SVDOverwrite SVDJob = 'O'
	SVDNone      SVDJob = 'N'
)

type GSVDJob byte

const (
	GSVDU    GSVDJob = 'U'
	GSVDV    GSVDJob = 'V'
	GSVDQ    GSVDJob = 'Q'
	GSVDUnit GSVDJob = 'I'
	GSVDNone GSVDJob = 'N'
)

type EVComp byte

const (
	EVOrig     EVComp = 'V'
	EVTridiag  EVComp = 'I'
	EVCompNone EVComp = 'N'
)

type EVJob byte

const (
	EVCompute EVJob = 'V'
	EVNone    EVJob = 'N'
)

type LeftEVJob byte

const (
	LeftEVCompute LeftEVJob = 'V'
	LeftEVNone    LeftEVJob = 'N'
)

type RightEVJob byte

const (
	RightEVCompute RightEVJob = 'V'
	RightEVNone    RightEVJob = 'N'
)

type BalanceJob byte

const (
	Permute      BalanceJob = 'P'
	Scale        BalanceJob = 'S'
	PermuteScale BalanceJob = 'B'
	BalanceNone  BalanceJob = 'N'
)

type SchurJob byte

const (
	EigenvaluesOnly     SchurJob = 'E'
	EigenvaluesAndSchur SchurJob = 'S'
)

type SchurComp byte

const (
	SchurOrig SchurComp = 'V'
	SchurHess SchurComp = 'I'
	SchurNone SchurComp = 'N'
)

type UpdateSchurComp byte

const (
	UpdateSchur     UpdateSchurComp = 'V'
	UpdateSchurNone UpdateSchurComp = 'N'
)

type EVSide byte

const (
	EVRight EVSide = 'R'
	EVLeft  EVSide = 'L'
	EVBoth  EVSide = 'B'
)

type EVHowMany byte

const (
	EVAll      EVHowMany = 'A'
	EVAllMulQ  EVHowMany = 'B'
	EVSelected EVHowMany = 'S'
)

type MaximizeNormXJob byte

const (
	LocalLookAhead       MaximizeNormXJob = 0
	NormalizedNullVector MaximizeNormXJob = 2
)

type OrthoComp byte

const (
	OrthoNone     OrthoComp = 'N'
	OrthoExplicit OrthoComp = 'I'
	OrthoPostmul  OrthoComp = 'V'
)
